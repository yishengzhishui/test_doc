## 变量

```go
var a int =10

//如果你没有显式为变量赋予初值，Go 编译器会为变量赋予这个类型的零值
var a int // a的初值为int类型的零值：0

//变量声明块
var (
    a int = 128
    b int8 = 6
    s string = "hello"
    c rune = 'A'
    t bool = true
)

var a, b, c int = 5, 6, 7

//默认类型赋值
var b = 13//整型值的默认类型 int
// 
var b = int32(13)  //显式地为变量指定类型

// 短变量声明
a := 12
b := 'A'
c := "hello"

a, b, c := 12, 'A', "hello"
```

Go 语言的两类变量。

一类称为**包级变量 (package varible)**，也就是在包级别可见的变量。如果是导出变量（大写字母开头），那么这个包级变量也可以被视为全局变量；**包级变量只能使用带有 var 关键字的变量声明形式**，不能使用短变量声明形式。

另一类则是**局部变量 (local varible)**，也就是 Go 函数或方法体内声明的变量，仅在函数或方法体内可见。



## 常量

```go
package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}
```

`const` 用于声明一个常量



## If/Else

在 Go 中，条件语句的圆括号不是必需的，但是花括号是必需的。

**在条件语句之前可以有一个声明语句；在这里声明的变量可以在这个语句所有的条件分支中使用**。



```go
package main

import "fmt"

func main() {

    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
```

## Switch

在同一个 `case` 语句中，你可以使用逗号来分隔多个表达式。 在这个例子中，我们还使用了可选的 `default` 分支。

```go
package main

import (
    "fmt"
    "time"
)

func main() {

    i := 2
    fmt.Print("write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
// print
Write 2 as two
It's a weekday
It's after noon
I'm a bool
I'm an int
Don't know type string
```



## 数组

*数组* 是一个具有编号且长度固定的元素序列

```go
package main

import "fmt"

func main() {

    var a [5]int //创建了一个刚好可以存放 5 个 int 元素的数组 a
    fmt.Println("emp:", a)

    a[4] = 100 //可以使用 array[index] = value 语法来设置数组指定位置的值， 或者用 array[index] 得到值
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5} //一行内声明并初始化一个数组
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
  

values := [5]int{1, 2, 3, 4, 5}
values[1:3] // {2, 3, } 不包含尾部
len(values) // 5
}
```



## 切片

*Slice* 是 Go 中一个重要的数据类型，它提供了比数组更强大的序列交互方式。

```go
package main

import "fmt"

func main() {

    s := make([]string, 3)//创建了一个长度为 3 的 string 类型的 slice（初始值为零值）
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s) //apd: [a b c d e f]

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c) //cpy: [a b c d e f]

    l := s[2:5]
    fmt.Println("sl1:", l) //包含元素 s[2]、s[3] 和 s[4] 的 slice

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

  //Slice 可以组成多维数据结构。内部的 slice 长度可以不一致，这一点和多维数组不同
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD) //2d:  [[0] [1 2] [2 3 4]]
  
  ss := make([]int, 1,2)
  ss2 := []int{1,3,4}
  // make function: create a slice with length and capacity
slice := make([]int, 5, 6) // make(type, len, cap)
  // For range: iterate over a slice
slice := string["W", "o", "w"]

for i, value := range slice {
    i // 0, then 1, then 2
    value // "W", then "o", then "w"
}

// Skip index or value

for i := range slice {
    i // 0, then 1, then 2
}

for _, value := range slice {
   value // "W", then "o", then "w"
}
}
```

## Map

*map* 是 Go 内建的[关联数据类型](http://zh.wikipedia.org/wiki/关联数组) （在一些其他的语言中也被称为 *哈希(hash)* 或者 *字典(dict)* ）。

```go
package main

import "fmt"

func main() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m) //使用 fmt.Println 打印一个 map，会输出它所有的键值对。

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)
//选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}  //声明并初始化
    fmt.Println("map:", n)
}
// result
map: map[k1:7 k2:13]
v1:  7
len: 2
map: map[k1:7]
prs: false
map: map[foo:1 bar:2]
```



## For循环

```go
package main

import "fmt"

func main() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
```

## Range 遍历

```go
package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

  //range 也可以只遍历 map 的键。
    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
  // range 在字符串中迭代 unicode 码点(code point)。 第一个返回值是字符的起始字节位置(相对的偏移量)，然后第二个是字符本身(码点值) 
  var s = "中国人"

for i, v := range s {
    fmt.Printf("index: %d, value: 0x%x\n", i, v)
}
//print
index: 0, value: 0x4e2d
index: 3, value: 0x56fd
index: 6, value: 0x4eba
}
// result
sum: 9
index: 1
a -> apple
b -> banana
key: a
key: b
0 103
1 111
```



## 函数：function

参数传递

```go
package main

import "fmt"

func plus(a int, b int) int {

    return a + b //需要明确的 return，它不会自动 return 最后一个表达式的值
}

//当多个连续的参数为同样类型时， 可以仅声明最后一个参数的类型
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {

    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
```

多返回值

```go
package main

import "fmt"

func vals() (int, int) {
    return 3, 7
}

func main() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals() //你仅仅需要返回值的一部分的话，你可以使用空白标识符 _
    fmt.Println(c)
}
```



可变参数函数

```go
package main

import "fmt"
//接受任意数量的 int 作为参数
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...) //多个值的 slice，想把它们作为参数使用
}
```

匿名函数和闭包

```go
package main

import "fmt"

// intSeq 函数返回一个在其函数体内定义的匿名函数。 返回的函数使用闭包的方式 隐藏 变量 i。 返回的函数 隐藏 变量 i 以形成闭包
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    nextInt := intSeq() // 获得了闭包函数

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
}
//result
1
2
3
1
```

递归

```go
package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1) // fact 函数在到达 fact(0) 前一直调用自身
}

func main() {
    fmt.Println(fact(7))

// 闭包也可以是递归的，但这要求在定义闭包之前用类型化的 var 显式声明闭包
    var fib func(n int) int 

    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }

    fmt.Println(fib(7))
}
```































