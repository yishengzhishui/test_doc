## 基础

### 项目调试运行异常，可能是IDE文件识别异常，重新下载就行了

### 调试配置需要调整模版，为package(软件包)

### main 函数特性

如果 main 函数里面引用了同一个包的其它方法、类型，那么要加上对应的文件，可以：

* go run 命令加上其它文件，例如 `go run main.go hello.go`
* `go run .`
* `go build .` 然后执行 `./hello_go`

### Package声明

每个文件必须有package声明

同一个目录下，go文件package 声明必须一致

同一个目录下，go文件和test.go文件package可以不一致可以

package名字可以和目录名字不一样

### 字符串使用注意

```go
// println("hello, " + 123) 不行
// //println("hello, " + string(123)) 也不行
println(fmt.Sprintf("hello, %d", 123))
```

字节数len()

字符个数与编码有关，默认是utf8：utf8.RuneCountInString()

```go
//字符个数
println(len("你好"))                      // 输出 6
println(utf8.RuneCountInString("你好"))   // 输出2
println(utf8.RuneCountInString("你好ab")) // 输出4
```

### iota

`iota` 是 Go 语言中的一个预定义标识符，用于在常量声明中生成连续的整数值。`iota` 在每次 `const` 关键字出现时被重置为0，然后在下一个 `const` 出现之前，每出现一次就自增1。

以下是一个简单的使用 `iota` 的示例：

```go
package main

import "fmt"

const (
    // 使用 iota 创建枚举
    Sunday = iota // 0
    Monday        // 1
    Tuesday       // 2
    Wednesday     // 3
    Thursday      // 4
    Friday        // 5
    Saturday      // 6
)

const (
    // 在同一 const 块中，iota 会继续自增
    A = iota // 0
    B        // 1
    C        // 2
)

const (
    // 可以通过空白标识符 (_) 跳过某些值
    D = iota // 0
    _        // 跳过 1
    F = iota // 2
)

func main() {
    // 输出常量值
    fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
    fmt.Println(A, B, C)
    fmt.Println(D, F)
}
```

在上述示例中：

- `Sunday` 到 `Saturday` 构成了一个星期的枚举，每个常量对应一个整数值，从0开始自增。
- `A`、`B`、`C` 在同一 `const` 块中，`iota` 会继续自增。
- 使用 `_` 可以跳过某些值，此时 `iota` 仍然会自增。

`iota` 可以用于生成一系列相关值的常量，枚举类型是 `iota` 常用的场景之一。它简化了常量声明中的重复工作。

### 函数式编程and闭包

```go
// Func3 带名字的返回值
func Func3(a int, b int) (str string, err error) {
	res := "hello"
	// 虽然带名字，但是我们并没有用
	return res, nil
}
func Func4() {
	myFunc3 := Func3
	_, _ = myFunc3(1, 2)
}

// Func6 的返回值是一个方法，
func Func6() func(name string) string {
	return func(name string) string {
		return "hello," + name
	}
}

func Closure(name string) func() string {

	// 返回的这个函数，就是一个闭包。
	// 它引用到了 Closure 这个方法的入参
	return func() string {
		return "hello, " + name
	}
}
func main() {
	sayHello := Func6()
	a := sayHello("大明")
	fa := Closure("test")
	println(a)
	println(fa())
}


```

这`Func6`和`Closure`两个函数都是返回闭包的例子，但它们之间有一些关键的区别：

1. **参数的不同：**

   - 在第一个例子中，`Func6` 函数返回的闭包接受一个字符串参数 `name`。
   - 在第二个例子中，`Closure` 函数返回的闭包不再接受参数，而是直接引用了外部函数 `Closure` 中的参数 `name`。
2. **闭包的使用场景：**

   - 第一个例子中的闭包可以多次调用，并且每次调用时可以传入不同的 `name` 参数。
   - 第二个例子中的闭包只能在外部函数 `Closure` 调用的上下文中使用，并且每个闭包都引用了 `Closure` 中相同的 `name` 参数。

以下是使用这两个函数的示例：

```go
// 第一个例子
myClosure1 := Func6()
result1 := myClosure1("Alice")
fmt.Println(result1) // 输出：hello, Alice

// 第二个例子
myClosure2 := Closure("Bob") // "Bob"其实还在下面的result2中使用，所以不会被销毁
result2 := myClosure2()
fmt.Println(result2) // 输出：hello, Bob
```

在第一个例子中，我们通过 `Func6` 获取一个闭包，并可以在多次调用中传递不同的参数。

在第二个例子中，我们通过 `Closure` 获取一个闭包，该闭包直接引用了 `name` 参数。在这种情况下，我们只能在 `Closure` 函数调用的上下文中使用闭包，并且每个闭包都引用了相同的 `name`。

总体而言，这两个例子展示了闭包如何引用外部变量，但它们的具体用途可能会因为参数的不同而有所不同。希望这能够帮助理解这两者之间的区别。

闭包如果使用不当可能会引起内存泄露。即一个对象被闭包引用的话，它是不会被垃圾回收的。

这句话的意思是，如果在闭包中引用了某个对象，而该闭包被持有（比如作为变量或返回值），那么该对象就不会被垃圾回收。这可能导致内存泄漏，因为即使在程序的其他地方不再使用这个对象，**由于闭包的存在，该对象仍然被保留在内存中**。

### defer

#### defer调用机制

`defer` 是 Go 语言中用于延迟执行函数调用的关键字。当使用 `defer` 时，函数调用不会立即执行，而是在包含 `defer` 语句的函数执行完毕后才执行。`defer` 的调用机制可以简要描述为以下几点：

1. **延迟执行：** 当一个函数中包含了 `defer` 语句时，被延迟执行的函数调用会在包含该语句的函数即将返回时执行，而不是在 `defer` 语句出现的位置立即执行。
2. **后进先出（LIFO）：** 如果一个函数中包含多个 `defer` 语句，它们会按照后进先出的顺序执行。也就是说，最后一个被推入（deferred）的函数调用会最先被执行，依此类推。

以下是一个简单的例子，用于说明 `defer` 的调用机制：

```go
package main

import "fmt"

func main() {
	fmt.Println("Start")

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")

	fmt.Println("End")
}
```

在这个例子中，输出的顺序是：

```
Start
End
Deferred 3
Deferred 2
Deferred 1
```

可以看到，被 `defer` 延迟执行的函数调用是在包含 `defer` 语句的函数即将返回时执行的，而且按照后进先出的顺序执行。

`defer` 的典型用途包括资源释放、文件关闭、解锁等操作，以确保在函数执行完毕之前执行这些清理工作。

```go
func DeferClosure() {
	i := 0
	defer func() {
		println(i) // 结果是1
	}()
	i = 1
}

//参数传递
func DeferClosureV1() {
	i := 0
	defer func(val int) {
		println(val) //结果是0
	}(i)
	i = 1
}
```

* 作为参数传入的:定义defer的时候就确定了。
* 作为闭包引入的:执行defer对应的方法的时候才确定。

#### defer修改返回值

![image.png](./assets/1700649557776-image.png)

```go
func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}
func DeferClosureLoopV4() {
		var j int
		for i := 0; i < 10; i++ {
			j = i
			defer func() {
				println(j)
			}()
		}
	}
func main() {
	DeferClosureLoopV1() // 都是10
	DeferClosureLoopV2() // 9～0
	DeferClosureLoopV3() // 9～0
        DeferClosureLoopV3() // 都是9
}
```

`DeferClosureLoopV1()`：在这个方法中，`defer`语句延迟执行了一个匿名函数，而这个匿名函数中**打印的`i`是对外部循环变量的引用**。由于`defer`语句是在函数返回时才执行的，当实际执行这些`defer`语句时，循环已经执行完毕，此时`i`的值已经是 10。因此，该方法会打印出 10 个`10`。

`DeferClosureLoopV2()`:在这个方法中，通过将循环变量的值作为参数传递给闭包，**避免了闭包捕获外部变量的问题**。每个 `defer`语句中的闭包都捕获了不同的`val`值，因此在执行时，会打印出从 9 到 0 的数字。这是因为`i` 的值在每次迭代时都被传递给了闭包，而不是引用外部的循环变量。

`DeferClosureLoopV3()`:在这个方法中，通过在每次循环迭代时创建一个新的变量 `j`来**避免闭包捕获外部变量**。每个`defer`语句中的闭包都引用了不同的`j` 变量，因此在执行时，同样会打印出从 9 到 0 的数字。

`DeferClosureLoopV4()`:在这个代码中，`defer` 语句中的闭包引用了外部的变量 `j`，而且 `j` 是在循环中不断被更新的。因为 `defer` 语句是在函数返回时才执行的，所以在实际执行时，`j` 的值已经是循环结束后的值，即 `j` 的最终值为 9。因此，无论循环迭代多少次，`defer` 语句中的闭包都会引用到最终的 `j` 值

### 方法调用总结

* Go 方法的作用域和变量作用域一样，通过大小写控制。
* Go 的返回值是可以有名字的，可以通过给予名字让调用方清楚知道你返回的是什么。
* Go 中方法是一等公民，所以函数式编程非常常见。在初学的时候，不需要掌握函数式编程，确保自己能够看得懂就可以。
* 闭包是指一个方法与跟着这个方法绑定的运行时刻上下文。初学的时候不要求掌握闭包用法，确保能看懂就行。面试的时候要能回答出来。
* defer 是先进后出，或者说后进先出。

### 子切片的共享内存

核心:共享数组。
子切片和切片究竟会不会互相影响，就抓住一点:它们是不是还共享数组?

* 就是如果它们结构没有变化，那肯定是共享的;
* 但是结构变化了，就可能不是共享了。

什么情况下结构会发生变化?**扩容了**。

所以，切片与子切片，切片作为参数传递到别的方法、结构体里面，任何情况下你要判断是否内存共享，那么就一个点:有没有扩容。

```go
func ShareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("share slice s1: %v len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("share slice s2: %v len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[0] = 99

	fmt.Printf("s2[0]=99 share slice s1: %v len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2[0]=99 share slice s2: %v len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 199)
	fmt.Printf("append s2 share slice s1: %v len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("append s2 share slice s2: %v len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[0] = 1999
	fmt.Printf("s2[0] = 1999 share slice s1: %v len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2[0] = 1999 share slice s2: %v len: %d, cap: %d \n", s2, len(s2), cap(s2))
}
// result
share slice s1: [1 2 3 4] len: 4, cap: 4 
share slice s2: [3 4] len: 2, cap: 2 
s2[0]=99 share slice s1: [1 2 99 4] len: 4, cap: 4 
s2[0]=99 share slice s2: [99 4] len: 2, cap: 2 
append s2 share slice s1: [1 2 99 4] len: 4, cap: 4 
append s2 share slice s2: [99 4 199] len: 3, cap: 4 
s2[0] = 1999 share slice s1: [1 2 99 4] len: 4, cap: 4 
s2[0] = 1999 share slice s2: [1999 4 199] len: 3, cap: 4 

```

### map

map 的遍历是随机的，也就是说你遍历两遍，输出的结果都不一样。

可以用delete删除， `delete(map,'key')`

### comparable

* 在switch里面，值必须是可比较的。
* 在map里面，key也必须是可比较的。

所谓可比较的(comparable)在 Go 里面就是指:Go 在编译的时候、运行的时候能够判断出来元素是不是相等。

• 基本类型和string都是可比较的。

• 如果元素是可比较的，那么该数组也是可比较的。

### 结构体

#### 结构体初始化-Go 没有构造函数

* 初始化语法:Struct{}
* 获取指针:`&Struct{}`、`new(Struct)`

new 可以理解为 Go 会为你的变量分配内存，并且把内存都置为0

如果声明了一个指针，但是没有赋值，那么是nil
