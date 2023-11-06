## 1.第一讲 并发编程

![image.png](./assets/1691675805006-image.png)

### Context

`context` 包是 Go 语言中用于处理请求范围数据、取消信号和截止时间的标准库。它提供了一种在跨 API 边界和进程边界传递请求范围数据的方式，同时支持取消信号和截止时间的传播。

**context实例是不可变的，每次都是新创建的**

在 Go 中，`Context` 实例是不可变的，这意味着一旦创建，就不能再修改。如果你需要在 `Context` 中传递新的值，你需要创建一个新的 `Context` 实例，以便在整个调用链上传递。

以下是 `context` 包的主要组件和概念：

1. **Context 接口：**

   - `Context` 是一个接口类型，定义了用于处理请求范围数据的方法。
   - 标准库提供了两个基础的 `Context` 实现，分别是 `context.Background()` 和 `context.TODO()`。
2. **`context.Background()`：**

   - `context.Background()` 返回一个空的、**非取消**的 `Context`。它通常用作根 `Context`。
3. **`context.TODO()`：**

   - `context.TODO()` 和 `context.Background()` 类似，但它表明代码中应该没有处理 `Context` 的具体逻辑。它通常在还没有明确的 `Context` 时使用。
4. **`context.WithCancel`：**

   - `context.WithCancel` 返回一个**带有取消函数**的新的 `Context`。调用取消函数会关闭该 `Context`，通知所有与之关联的 goroutine 取消操作。
   - ```go
     ctx, cancel := context.WithCancel(context.Background())
     defer cancel() // 确保在不再需要时取消
     ```
5. **`context.WithTimeout` 和 `context.WithDeadline`（本质上是一样的）：**

   - `context.WithTimeout` 返回一个在**超时时间到达时自动取消**的 `Context`。
   - `context.WithDeadline` 返回一个在**指定截止时间到达时自动取消**的 `Context`。
   - ```go
     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
     dlCtx, cancel2 := context.WithDeadline(ctx, time.Now().Add(time.Minute))
     defer cancel2()
     defer cancel() // 确保在不再需要时取消
     ```
6. **`context.Value`：**

   - **`context.WithValue(parent, key, value)`：** 返回一个带有键值对的新 `Context`，用于传递请求范围的数据。这是一种将元数据传递到整个调用链的方式。
   - ```go
     type key int

     const myKey key = 42

     ctx := context.WithValue(context.Background(), myKey, "some value")
     ```

使用 `context` 的主要场景包括：

- 在 HTTP 请求处理中传递取消信号和截止时间。
- 在多个 goroutine 之间传递请求范围的数据。

示例：

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带有取消函数的 Context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 通常通过 defer 来调用取消函数

	// 启动一个 goroutine，在取消信号到达时退出
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine: Context canceled")
		}
	}()

	// 模拟一些工作
	time.Sleep(2 * time.Second)

	// 发送取消信号
	cancel()
	time.Sleep(1 * time.Second) // 为了确保 Goroutine 有足够的时间响应
}
```

在这个示例中，`context.WithCancel` 用于创建一个带有取消函数的 `Context`，并在稍后调用 `cancel()` 时发送取消信号。 Goroutine 通过监听 `ctx.Done()` 通道，在取消信号到达时退出。

### 父子关系

context实例之间存在父子关系：

* 父节点取消或超时，所有派生的子节点全部取消或超时
* 当找key的时候，子context先查找自己，找不到就到上一级找

**总结一下就是：控制是从上至下的，查找是从下至上的**

另外 子context重置超时时间是不会生效的，依旧收到父context的控制

一般父节点是无法访问子context的内容的，如果逼不得已，可以采用在父context中放入一个map，后续都修改map的方法得到子context数据

```go
func TestParentValueCtx(t *testing.T) {
	ctx := context.Background()
	childCtx := context.WithValue(ctx, "map", map[string]string{})
	ccChild := context.WithValue(childCtx, "key1", "value1")
	m := ccChild.Value("map").(map[string]string)
	fmt.Println(m)
	m["key1"] = "val1"
	val := childCtx.Value("key1")
	fmt.Println(val)
	val = childCtx.Value("map")
	fmt.Println(val)
}
```

### 简单例子：

#### 超时控制

```go
func MyBusiness() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("hello, world")
}
func TestBusinessTimeout(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	end := make(chan struct{}, 1)
	go func() {
		MyBusiness()
		end <- struct{}{}
	}()
	ch := timeoutCtx.Done()
	select {
	case <-ch:
		fmt.Println("timeout")
	case <-end:
		fmt.Println("business end")
	}
}

// 结果
hello， world
business end
```

另一种超时控制（其实是定时任务）不是很推荐

`time.AfterFunc` 有两个问题：就是如果不主动取消，超过多长时间必然会执行的；另外有一个问题就是如果主动取消，可能是存在一个短暂的时间差

```go
func TestTimeoutTimeAfter(t *testing.T) {
	bsChan := make(chan struct{})
	go func() {
		slowBusiness()
		bsChan <- struct{}{}
	}()

	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("timeout")
	})
	<-bsChan
	fmt.Println("business end")
	timer.Stop()
}
```

#### 赋值拿取

```go
func TestContext(t *testing.T) {
	ctx := context.Background()
	valCtx := context.WithValue(ctx, "abc", 123)
	val := valCtx.Value("abc")
	fmt.Println(val)
}
```

### ErrGroup

`errgroup` 是 Go 语言标准库 `golang.org/x/sync/errgroup` 包中的一个类型，用于简化 goroutine 的错误处理。它提供了一种方便的方式来等待一组 goroutine 完成，并能够在其中任何一个返回错误时取消整个组的执行。

以下是 `errgroup` 的简要用法：

1. 导入 `golang.org/x/sync/errgroup` 包。

```go
import "golang.org/x/sync/errgroup"
```

2. 创建一个 `errgroup.Group` 对象。

```go
var g errgroup.Group
```

3. 启动 goroutine，并使用 `g.Go` 方法添加到组中。

```go
g.Go(func() error {
    // Your goroutine logic here
    return nil // or an error
})
```

4. 使用 `g.Wait()` 来等待所有 goroutine 完成。如果其中任何一个返回非空的错误，`Wait` 会立即取消所有其他 goroutine 的执行。

```go
if err := g.Wait(); err != nil {
    // Handle the error
}
```

这样，`errgroup` 可以很方便地管理一组 goroutine 的执行，如果其中任何一个返回错误，它将会取消其他所有正在执行的 goroutine，并将该错误返回给调用者。

以下是一个简单的示例，演示了 `errgroup` 的基本用法：

```go
package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group

	// Goroutine 1
	g.Go(func() error {
		fmt.Println("Goroutine 1")
		return nil
	})

	// Goroutine 2
	g.Go(func() error {
		fmt.Println("Goroutine 2")
		return fmt.Errorf("An error occurred in Goroutine 2")
	})

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	}
}
```

在这个例子中，尽管 Goroutine 2 返回了一个错误，但由于使用了 `errgroup`，它会导致所有 goroutine 立即停止执行，最后的 `Wait` 会捕获并输出错误。

示例2：

errgroup 使用context

```go
func TestErrgroup(t *testing.T) {
//errgroup.WithContext 创建一个与 context.Background() 关联的新 errgroup.Group，这样就可以使用 context.Context 来取消所有任务。
	eg, ctx := errgroup.WithContext(context.Background())
	var result int64 = 0
	for i := 0; i < 10; i++ {
		delta := i
		eg.Go(func() error {
//使用 atomic.AddInt64 来原子地将 delta 添加到 result 中
			atomic.AddInt64(&result, int64(delta))
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
//虽然这里调用了 ctx.Err()，但是没有对其返回值进行处理。通常，你应该检查 ctx.Err() 是否为 context.Canceled 或 context.DeadlineExceeded，以确定是任务正常完成还是由于上下文取消或超时而退出。
	ctx.Err()
	fmt.Println(result)
}
```

`errgroup` 与 `context` 结合的主要应用场景是管理多个并发任务，同时能够及时地处理其中的错误，实现任务的同步和取消。

虽然 `errgroup` 在某个 goroutine 出现错误时会取消其他任务，但 `context` 主要用于提供对整个任务组的上下文管理，包括超时控制和手动取消。

1. **超时控制：** 使用 `context` 可以方便地为整个任务组设置超时。`context` 提供了 `WithTimeout` 和 `WithDeadline` 等方法，使得你可以在一定时间内完成一组任务，防止任务执行时间过长。
2. **手动取消：** 在某些情况下，你可能需要手动取消一组任务，而不是等待其中一个任务返回错误。使用 `context` 提供的 `WithCancel` 方法，你可以在需要的时候手动取消整个任务组。
3. **上下文传递：** `context` 不仅用于管理整个任务组的上下文，还可以传递一些共享的上下文信息给每个任务。

   ```go
   parentCtx := context.Background()
   parentCtx = context.WithValue(parentCtx, "key", "value")

   eg, ctx := errgroup.WithContext(parentCtx)
   ```
4. **更多控制：** `context` 提供了更多的控制，例如 `Done()` 方法用于接收任务组的完成信号，以及 `Err()` 方法用于获取取消的原因。

总体来说，`context` 和 `errgroup` 结合使用，可以提供更全面的控制和管理，并更好地适应不同的需求场景。使用 `context` 可以使得任务组的管理更加灵活，更容易适应不同的场景和需求。

### 源码相关

#### cancelCtx

在 Go 中，`context` 包中的 `WithCancel` 函数创建了一个新的 `context.Context`，同时返回一个用于取消的 `cancel` 函数。这个 `cancel` 函数内部会关闭一个 `cancelCtx`，该 `cancelCtx` 是 `cancelCtx` 类型的结构体。这个结构体定义如下：

```go
type cancelCtx struct {
	Context
	done chan struct{} // 用于通知已经取消的信号
	err  error          // 存储取消的原因
	mu   sync.Mutex     // 保护以下字段
}

func (c *cancelCtx) Done() <-chan struct{} {
	return c.done
}

func (c *cancelCtx) Err() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.err
}

func (c *cancelCtx) String() string {
	return fmt.Sprintf("%v.WithCancel", c.Context)
}
```

这个结构体实现了 `Context` 接口，其中包含一个 `done` 通道用于通知已经取消的信号，一个 `err` 字段用于存储取消的原因，以及一个互斥锁 `mu` 用于保护这些字段。

`WithCancel` 函数创建了一个 `cancelCtx` 实例，将其作为新的 `Context` 返回。当调用 `cancel` 函数时，它会向 `done` 通道发送信号，表示 `Context` 已经被取消，并设置 `err` 字段为 `context.Canceled`。这样，任何监听 `Done` 通道的 goroutine 都会收到信号，知道 `Context` 已经取消。

在实际使用中，我们可以通过调用 `cancel` 函数来手动取消 `Context`，以便通知其相关的任务停止执行。

#### valueCtx和timerCtx

在 `context` 包中，`valueCtx` 和 `timeCtx` 是 `context` 包中两个具体的类型。这两个类型分别用于实现 `WithValue` 和 `WithTimeout` 函数。

1. **valueCtx**：`valueCtx` 实现了 `Context` 接口，它的主要目的是在 `WithValue` 函数中创建一个带有新值的上下文。它包含了一个 `key` 和一个 `value`，这对应于键值对，允许我们在上下文中存储和检索值。当你调用 `WithValue` 时，它会创建一个新的 `valueCtx`，并将这个键值对添加到上下文中。

   简化的伪代码如下：

   ```go
   type valueCtx struct {
       Context
       key, val interface{}
   }

   func WithValue(parent Context, key, val interface{}) Context {
       return &valueCtx{
           Context: parent,
           key:     key,
           val:     val,
       }
   }
   ```
2. timerCtx

   `timerCtx` 是 `context` 包中的一个内部类型，用于实现具有超时功能的上下文。它的原理涉及 Go 语言的协程和定时器机制。

   1. **继承关系：** `timerCtx` 继承自 `cancelCtx`，这是因为超时是可以被取消的，而 `cancelCtx` 实现了 `Context` 接口。
   2. **定时器：** `timerCtx` 内部包含一个 `time.Timer` 对象，用于实现超时机制。`time.Timer` 是 Go 语言标准库中的一个定时器，它可以在指定的时间后发送一个事件给一个通道。
   3. **实现超时：** 当创建一个 `timerCtx` 时，会同时创建一个 `time.Timer`，并将其保存在 `timerCtx` 中。这个定时器的超时时间就是上下文的超时时间。
   4. **取消：** 如果上下文的 `cancel` 方法被调用（例如通过调用 `cancel` 函数或者超时发生），`cancelCtx` 的 `cancel` 方法会被触发。这个方法会关闭与父上下文的连接，并取消子上下文。同时，如果有定时器在运行，它也会被停止。
   5. **协程：** 在 Go 中，协程是轻量级的线程。当一个协程（goroutine）使用了 `timerCtx` 创建的上下文，并且超时时间到达，`timerCtx` 会取消它的父上下文，并通过关闭 `timerCtx.Done` 通道通知相关的协程。

   这样，通过 `timerCtx` 的机制，我们可以在一段时间后自动取消与上下文相关联的任务。这对于处理超时场景非常有用，例如在网络请求中设置超时时间，确保不会因为等待太久而导致整个程序变得不响应。

   **简单的回答**：在 Go 语言中，`timerCtx` 是 `context` 包中的一个实现，用于在一段时间后自动取消与上下文相关联的任务。它基于 `time.Timer` 实现了超时机制，当上下文的超时时间到达时，`timerCtx` 会触发取消操作，并通过关闭 `Done` 通道通知相关的协程。这使得我们能够很方便地处理一些需要在一定时间内完成的任务，比如设置网络请求的超时时间。

### 注意事项：

* 一般只用作**方法参数**，而且是作为第一个参数
* 所有的公共方法，除非是util，helper之类的，否则**都加上context参数**
* **不要用作结构体字段**，除非把你的结构体本身也是表达一个上下文的概念（http的request）

#### 面试要点

* context.Context 使用场景：上下文传递与超时控制
* context.Context 原理：父亲如何控制儿子：通过儿子主动加入到父亲的children里面，父亲只要遍历就行
* valueCtx和timeCtx的原理
