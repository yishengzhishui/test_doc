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
5. **`context.WithTimeout` 和 `context.WithDeadline`：**

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
