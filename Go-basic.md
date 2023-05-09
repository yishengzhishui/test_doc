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

一类称为**包级变量 (package varible)**，也就是在包级别可见的变量。如果是导出变量（大写字母开头），那么这个包级变量也可以被视为全局变量；**包级变量只能使用带有 var 关键字的变量声明形式**
，不能使用短变量声明形式。

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

## 指针

Go 拥有指针。指针保存了值的**内存地址**。

类型 `*T` 是指向 `T` 类型值的指针。其零值为 `nil`，如果我们拥有一个类型 T，那么以 T 作为基类型的指针类型为 *T。

```go
var p *int
var a *T
```

`&` 操作符会生成一个指向其操作数的指针。&符号称为取地址符号

```
i := 42
p = &i
```

`*` 操作符表示指针指向的底层值。通过指针变量读取或修改其指向的内存地址上的变量值，这个操作被称为指针的解引用（dereference）。

```
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
```

这也就是通常所说的“间接引用”或“重定向”。

```go
package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

//参数为int指针
func zeroptr(iptr *int) {
    *iptr = 0 //通过指针设置值
}

func main() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
  
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}
// result
42
21
73

initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100


```

### 二级指针

```go

package main

func main() {
    var a int = 5
    var p1 *int = &a
    println(*p1) // 5
    var b int = 55
    var p2 *int = &b
    println(*p2) // 55

    var pp **int = &p1
    println(**pp) // 5
    pp = &p2      
    println(**pp) // 55
}  
------------------

package main

func foo(pp **int) {
    var b int = 55
    var p1 *int = &b
    (*pp) = p1
}

func main() {
    var a int = 5
    var p *int = &a
    println(*p) // 5
    foo(&p)
    println(*p) // 55
}
```

## 字符串和rune类型

Go语言中的字符串是一个只读的byte类型的切片。在Go语言当中，字符的概念被称为 `rune` - 它是一个表示 Unicode 编码的整数。

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {

    const s = "สวัสดี"

    fmt.Println("Len:", len(s))

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        examineRune(runeValue)
    }
}

func examineRune(r rune) {

    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}

Len: 18
e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
Rune count: 6
U+0E2A 'ส' starts at 0
U+0E27 'ว' starts at 3
U+0E31 'ั' starts at 6
U+0E2A 'ส' starts at 9
U+0E14 'ด' starts at 12
U+0E35 'ี' starts at 15
Using DecodeRuneInString
U+0E2A 'ส' starts at 0
found so sua
U+0E27 'ว' starts at 3
U+0E31 'ั' starts at 6
U+0E2A 'ส' starts at 9
found so sua
U+0E14 'ด' starts at 12
U+0E35 'ี' starts at 15

```

1)string 类型的数据是不可变的，提高了字符串的并发安全性和存储利用率。Go 语言规定，字符串类型的值在它的生命周期内是不可改变的。

```go
var s string = "hello"
s[0] = 'k'   // 错误：字符串的内容是不可改变的
s = "gopher" // ok
```

2)原生支持“所见即所得”的原始字符串，大大降低构造多行字符串时的心智负担

通过一对反引号原生支持构造“所见即所得”的原始字符串（Raw String）。而且，Go 语言原始字符串中的任意转义字符都不会起到转义的作用

```go
var s string = `         ,_---~~~~~----._
    _,,_,*^____      _____*g*\"*,--,
   / __/ /'     ^.  /      \ ^@q   f
  [  @f | @))    |  | @))   l  0 _/
   \/   \~____ / __ \_____/     \
    |           _l__l_           I
    }          [______]           I
    ]            | | |            |
    ]             ~ ~             |
    |                            |
     |                           |`
fmt.Println(s)
```

3)对非 ASCII 字符提供原生支持，消除了源码在不同环境下显示乱码的可能。Go 语言源文件默认采用的是 Unicode 字符集，

Go 字符串的组成

Go 语言在看待 Go 字符串组成这个问题上，有两种视角。

一种是**字节视角**，也就是和所有其它支持字符串的主流语言一样，Go 语言中的字符串值也是一个**可空的字节序列**，字节序列中的字节个数称为该字符串的长度。一个个的字节只是孤立数据，不表意。

```go
var s = "中国人"
fmt.Printf("the length of s = %d\n", len(s)) // 9

for i := 0; i < len(s); i++ {
  fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
}
fmt.Printf("\n")
```

另外一个**字符视角**，也就是字符串是由一个可空的字符序列构成。这个时候我们再看下面代码，以 0x4e2d 为例，它是汉字“中”在 Unicode 字符集表中的码点（Code Point）。

```go
var s = "中国人"
fmt.Println("the character count in s is", utf8.RuneCountInString(s)) // 3

for _, c := range s {
  fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba
}
fmt.Printf("\n")
```

### 码点

Unicode 字符集中的每个字符，都被分配了统一且唯一的字符编号。所谓 Unicode 码点，就是指将 Unicode 字符集中的所有字符“排成一队”，字符在这个“队伍”中的位次，就是它在 Unicode
字符集中的码点。也就说，一个码点唯一对应一个字符。“码点”的概念和我们马上要讲的 rune 类型有很大关系。

#### rune 类型与字符字面值

Go 使用 rune 这个类型来表示一个 Unicode 码点。rune 本质上是 int32 类型的别名类型，它与 int32 类型是完全等价的。

由于一个 Unicode 码点唯一对应一个 Unicode 字符。所以我们可以说，**一个 rune 实例就是一个 Unicode 字符，一个 Go 字符串也可以被视为 rune 实例的集合**。我们可以通过字符字面值来初始化一个 rune
变量。

```go
"abc\n"
"中国人"
"\u4e2d\u56fd\u4eba" // 中国人
"\U00004e2d\U000056fd\U00004eba" // 中国人
"中\u56fd\u4eba" // 中国人，不同字符字面值形式混合在一起
"\xe4\xb8\xad\xe5\x9b\xbd\xe4\xba\xba" // 十六进制表示的字符串字面值：中国人。  这个字节序列实际上是“中国人”这个 Unicode 字符串的 UTF-8 编码值
```

#### UTF-8 编码方案

UTF-8 编码解决的是 Unicode 码点值在计算机中如何存储和表示（位模式）的问题。那你可能会说，码点唯一确定一个 Unicode 字符，直接用码点值不行么？

和 UTF-32 方案不同，UTF-8 方案使用变长度字节，对 Unicode 字符的码点进行编码。编码采用的字节数量与 Unicode
字符在码点表中的序号有关：表示序号（码点）小的字符使用的字节数量少，表示序号（码点）大的字符使用的字节数多。

#### Go 字符串类型的内部表示

**string 类型其实是一个“描述符”，它本身并不真正存储字符串数据，而仅是由一个指向底层存储的指针和字符串的长度字段组成的**

![img](/Users/wangxing/Desktop/go/Go基础.assets/6c94a2f5a0f942e361792b26f5abfa28.jpg)

了解了 string 类型的实现原理后，我们还可以得到这样一个结论，那就是我们直接将 string 类型通过函数 / 方法参数传入也不会带来太多的开销。**因为传入的仅仅是一个“描述符”**，而不是真正的字符串数据。

#### Go 字符串类型的常见操作

由于字符串的不可变性，针对字符串，我们更多是尝试对其进行读取

##### 第一个操作：下标操作。

在字符串的实现中，真正存储数据的是底层的数组。字符串的下标操作本质上等价于底层数组的下标操作。

```go
var s = "中国人"
fmt.Printf("0x%x\n", s[0]) // 0xe4：字符“中” utf-8编码的第一个字节
```

我们可以看到，通过下标操作，我们获取的是字符串中**特定下标上的字节**，而不是字符

##### 第二个操作：字符迭代。

Go 有两种迭代形式：常规 for 迭代与 for range 迭代。

通过常规 **for 迭代**对字符串进行的操作是一种**字节视角**的迭代，每轮迭代得到的的结果都是组成字符串内容的一个字节，以及该字节所在的下标值，这也等价于对字符串底层数组的迭代，比如下面代码：

```go
var s = "中国人"

for i := 0; i < len(s); i++ {
  fmt.Printf("index: %d, value: 0x%x\n", i, s[i])
}
//print
index: 0, value: 0xe4
index: 1, value: 0xb8
index: 2, value: 0xad
index: 3, value: 0xe5
index: 4, value: 0x9b
index: 5, value: 0xbd
index: 6, value: 0xe4
index: 7, value: 0xba
index: 8, value: 0xba
```

for range 迭代

```go
var s = "中国人"

for i, v := range s {
    fmt.Printf("index: %d, value: 0x%x\n", i, v)
}
//print
index: 0, value: 0x4e2d
index: 3, value: 0x56fd
index: 6, value: 0x4eba
```

我们看到，通过 for range 迭代，我们每轮迭代得到的是字符串中 Unicode 字符的码点值，以及该字符在字符串中的偏移值。

##### 第三个操作：字符串连接

```go
s := "Rob Pike, "
s = s + "Robert Griesemer, "
s += " Ken Thompson"

fmt.Println(s) // Rob Pike, Robert Griesemer, Ken Thompson
```

虽然通过 +/+= 进行字符串连接的开发体验是最好的，但连接性能就未必是最快的了。除了这个方法外，Go 还提供了 strings.Builder、strings.Join、fmt.Sprintf 等函数来进行字符串连接操作。

> 如果能知道拼接字符串的个数，那么使用bytes.Buffer和strings.Builder的Grows申请空间后，性能是最好的；
>
> 如果不能确定长度，那么bytes.Buffer和strings.Builder也比“+”和fmt.Sprintf性能好很多。

> bytes.Buffer与strings.Builder，strings.Builder更合适，因为bytes.Buffer 转化为字符串时重新申请了一块空间，存放生成的字符串变量，而 strings.Builder 直接将底层的 []byte 转换成了字符串类型返回了回来。
>
>  bytes.Buffer 的注释中还特意提到了： To build strings more efficiently, see the strings.Builder type.

##### 第四个操作：字符串比较

Go 采用字典序的比较策略，分别从**每个字符串的起始处**，开始逐个**字节**地对两个字符串类型变量进行比较。

当两个字符串之间出现了第一个不相同的元素，比较就结束了，这两个元素的比较结果就会做为串最终的比较结果。如果出现两个**字符串长度不同**的情况，长度比较小的字符串会**用空元素补**齐，空元素比其他非空元素都小。

```go
func main() {
        // ==
        s1 := "世界和平"
        s2 := "世界" + "和平"
        fmt.Println(s1 == s2) // true

        // !=
        s1 = "Go"
        s2 = "C"
        fmt.Println(s1 != s2) // true

        // < and <=
        s1 = "12345"
        s2 = "23456"
        fmt.Println(s1 < s2)  // true
        fmt.Println(s1 <= s2) // true

        // > and >=
        s1 = "12345"
        s2 = "123"
        fmt.Println(s1 > s2)  // true
        fmt.Println(s1 >= s2) // true
}
```

## 结构体（struct）

Go 的*结构体(struct)* 是带类型的字段(fields)集合。 这在组织数据时非常有用。

```go
package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}

func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"}) //省略的字段将被初始化为零值

    fmt.Println(&person{name: "Ann", age: 40}) //& 前缀生成一个结构体指针

    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name) // Sean

    sp := &s
    fmt.Println(sp.age) //50    指针会被自动解引用

    sp.age = 51
    fmt.Println(sp.age)
}
// result
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
```

### 结构体方法

使用指针接收者的原因有二：

首先，方法能够修改其接收者指向的值。

其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

值/指针接收器都可以用值或者指针调用

```go
package main

import "fmt"

type rect struct {
    width, height int
}

//指针接收器，可以直接修改结构体的值
func (r *rect) area() int {
    return r.width * r.height
}
// 值接收器 对原结构体的副本（拷贝）操作
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
//result
area:  50
perim: 30
area:  50
perim: 30

------------------------------
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())   // 50
}
```

## 接口

方法签名的集合叫做：_接口(Interfaces)_。

```go
package main

import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}
//如果一个变量实现了某个接口，我们就可以调用指定接口中的方法。 这儿有一个通用的 measure 函数，我们可以通过它来使用所有的 geometry。
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
//结构体类型 circle 和 rect 都实现了 geometry 接口， 所以我们可以将其实例作为 measure 的参数
    measure(r)
    measure(c)
}

------------------
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
// result
(&{Hello}, *main.T)
Hello
(3.141592653589793, main.F)
3.141592653589793
```

## Embedding 嵌入

Go支持对于结构体(struct)和接口(interfaces)的 *嵌入(embedding)* 以表达一种更加无缝的 *组合(composition)* 类型

```go
package main

import "fmt"

type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}
//一个 container 嵌入 了一个 base. 一个嵌入看起来像一个没有名字的字段
type container struct {
    base
    str string
}

func main() {
//当创建含有嵌入的结构体，必须对嵌入进行显式的初始化； 在这里使用嵌入的类型当作字段的名字
    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }

    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

    fmt.Println("also num:", co.base.num)
//由于 container 嵌入了 base，因此base的方法 也成为了 container 的方法。在这里我们直接在 co 上 调用了一个从 base 嵌入的方法。
    fmt.Println("describe:", co.describe())

    type describer interface {
        describe() string
    }
//  可以使用带有方法的嵌入结构来赋予接口实现到其他结构上。 
//  因为嵌入了 base ，在这里我们看到 container 也实现了 describer 接口。
    var d describer = co
    fmt.Println("describer:", d.describe())
}

// result
co={num: 1, str: some name}
also num: 1
describe: base with num=1
describer: base with num=1
```

## 错误处理

在 Go 语言中，如果一个类型实现了 Error() 方法，那么当该类型的实例作为错误值被返回时，将自动调用其 Error() 方法来生成错误信息的字符串表示。

```go
package main

import (
    "errors"
    "fmt"
)
## 在函数 f1 中，如果参数 arg 的值等于 42，则返回一个错误对象。
## 这里使用了 Go 语言内置的 errors.New 函数来创建一个新的错误对象。
func f1(arg int) (int, error) {
    if arg == 42 {

        return -1, errors.New("can't work with 42")

    }

    return arg + 3, nil
}

type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {

        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {

    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

# 类型断言e.(*argError)将错误值转换为类型为*argError的指针。
# 如果转换成功，则变量ae保存了这个指针，同时变量ok的值为true。
#在这种情况下，代码会输出ae.arg和ae.prob两个属性的值，即错误值中记录的参数值和错误信息。
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
```

在这个例子中，我们定义了一个自定义的错误类型 argError，并在其上实现了 Error() 方法，用于将错误信息格式化为一个字符串。 当在 f2 函数中返回该类型的指针时，由于它实现了 Error() 方法，所以在
fmt.Println() 函数打印错误信息时，会自动调用其 Error() 方法来生成错误信息字符串。

`argError`类型和`*argError`类型 有啥区别
`argError`和`*argError`分别表示不同的类型。`argError`是一个结构体类型，它包含了两个字段arg和prob。而`*argError`表示一个指向argError类型值的指针类型。

具体来说，`argError`类型表示的是一个值，这个值是argError结构体的实例，而`*argError`类型表示的是一个指针，这个指针指向一个argError结构体的实例。

Error() 函数是 `*argError` 类型的指针接收器方法。在 Go 中，一个类型的指针可以调用该类型的所有方法，同时也可以调用该类型的指针接收器方法。因此，在这个例子中，`*argError` 类型的指针可以调用 Error()
方法，而不需要使用 & 运算符来解引用指针。

## 类型断言

Go的类型断言是指在运行时动态地判断一个接口值所持有的值的类型是否为某个特定的类型。类型断言的语法形式如下：
`x.(T)`
其中，x 是一个接口值，T 是一个类型。如果 x 所持有的值的类型是 T，则类型断言返回 x 中的值和 true，否则返回零值和 false。

## 协程(goroutine)

协程(goroutine) 是轻量级的执行线程。 在 Go 中，协程（Goroutine）是轻量级的执行线程，它由 Go 运行时（Go runtime）管理。协程的执行类似于线程，但它们的创建和销毁要比线程更快，并且它们占用的内存更少。

要创建一个协程，在 Go 中只需要在函数调用前添加 go 关键字即可

```go
package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")
}
```

这段代码创建了两个协程（goroutine），一个在主函数中通过go关键字调用f()函数创建，另一个是通过匿名函数创建。f()
函数是一个简单的循环，打印出三个数值。匿名函数只打印一个字符串。由于协程运行在独立的线程中，因此它们可能不按照特定的顺序运行。 最后，使用time.Sleep()函数使主线程休眠一秒钟，以确保协程有足够的时间完成执行。

## 通道(Channels)

通道(channels) 是连接多个协程的管道。 你可以从一个协程将值发送到通道，然后在另一个协程中接收。

```go
package main

import "fmt"

func main() {
// 使用 make(chan val-type) 创建一个新的通道。 通道类型就是他们需要传递值的类型。
    messages := make(chan string)
// 使用 channel <- 语法 发送 一个新的值到通道中。 这里我们在一个新的协程中发送 "ping" 到上面创建的 messages 通道中。
    go func() { messages <- "ping" }()
// 使用 <-channel 语法从通道中 接收 一个值。 这里我们会收到在上面发送的 "ping" 消息并将其打印出来。
    msg := <-messages
    fmt.Println(msg)
}
```

默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。 这个特性允许我们，不使用任何其它的同步操作， 就可以在程序结尾处等待消息 "ping"。

### 通道缓冲

默认情况下，通道是**无缓冲**的，这意味着只有对应的接收（<- chan） 通道准备好接收时，才允许进行发送（chan <-）。 有缓冲通道 允许在没有**对应接收者**的情况下，缓存一定数量的值。
在这种情况下，发送者可以发送到通道而不会被阻塞，直到缓冲区填满。同样地，接收者可以从通道接收数据而不会被阻塞，直到缓冲区为空。

```go
package main

import "fmt"

func main() {
//make 了一个字符串通道，最多允许缓存 2 个值
    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```

在使用通道时，通常需要将它们与goroutine一起使用，以便可以在不同的goroutine之间发送和接收数据。可以在goroutine内部使用select语句，从多个通道接收数据。

```go
select {
case <- ch1:
    // 接收来自ch1的数据
case <- ch2:
    // 接收来自ch2的数据
default:
    // 没有通道准备好，执行默认操作
}
```

select语句可以等待多个通道中的任何一个准备好接收数据。如果没有通道准备好，select语句就会等待。可以使用default语句指定在没有通道准备好时要执行的默认操作。

### 通道方向

当使用通道作为函数的参数时，你可以指定这个通道是否为只读或只写。 该特性可以提升程序的类型安全。 通道可以是单向的或双向的。单向通道限制了通道的操作方向，即只能发送或只能接收。

```go
package main

import "fmt"
//ping 函数定义了一个只能发送数据的（只写）通道。 
//尝试从这个通道接收数据会是一个编译时错误。
func ping(pings chan<- string, msg string) {
    pings <- msg
}
//pong 函数接收两个通道，
// pings 仅用于接收数据（只读），pongs 仅用于发送数据（只写）。
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
```

### 超时操作

```go
package main

import (
    "fmt"
    "time"
)

func main() {

//这里是使用 select 实现一个超时操作。
//res := <- c1 等待结果，<-time.After 等待超时（1秒钟）以后发送的值。
//由于 select 默认处理第一个已准备好的接收操作，
//因此如果操作耗时超过了允许的 1 秒的话，将会执行超时 case。
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }
    
  //如果我们允许一个长一点的超时时间：3 秒， 就可以成功的从 c2 接收到值，并且打印出结果。  
    
    

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}
```

### 关闭通道

```go
package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}
```

通道关闭后，再执行拿取操作，将会立即得到一个零值，并返回一个对应的可选的布尔值 false 表示通道已关闭。

## Timer 定时器

定时器表示在未来某一时刻的独立事件。 你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道 如果你需要的仅仅是单纯的等待，使用 time.Sleep 就够了。 使用定时器的原因之一就是，你可以在定时器触发之前将其取消

```go
package main

import (
    "fmt"
    "time"
)

func main() {
//这里的定时器将等待 2 秒。
    timer1 := time.NewTimer(2 * time.Second)
//<-timer1.C 会一直阻塞， 直到定时器的通道 C 明确的发送了定时器失效的值。
    <-timer1.C
    fmt.Println("Timer 1 fired")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    //停止了定时器 因为此时 timer2 尚未到期，所以该方法返回 true。
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    time.Sleep(2 * time.Second)
}
```

## Ticker 打点器

打点器可以和定时器一样被停止。 打点器一旦停止，将不能再从它的通道中接收到值。 我们将在运行 1600ms 后停止这个打点器。

```go
package main

import (
    "fmt"
    "time"
)

func main() {

    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}
```

## 字符串处理函数

### 字符串基本操作函数：

1. len(str) int：返回字符串的长度。
2. +或fmt.Sprintf(str1, str2) string：字符串拼接。
3. strings.Split(str, sep) []string：字符串分割。
4. strings.Contains(str, substr) bool：判断字符串是否包含某个子串。
5. strings.HasPrefix(str, prefix) bool：判断字符串是否以某个前缀开头。
6. strings.HasSuffix(str, suffix) bool：判断字符串是否以某个后缀结尾。
7. strings.Index(str, substr) int：返回子串在字符串中第一次出现的位置，若不存在则返回-1。
8. strings.LastIndex(str, substr) int：返回子串在字符串中最后一次出现的位置，若不存在则返回-1。

### 字符串处理函数：

1. strings.ToUpper(str) string：将字符串中的所有字符转换为大写。
2. strings.ToLower(str) string：将字符串中的所有字符转换为小写。 
3. strings.TrimSpace(str) string：去除字符串首尾的空白字符。 
4. strings.Trim(str, cutset) string：去除字符串首尾指定的字符。 
5. strings.Replace(str, old, new, n) string：将字符串中的某个子串替换成另一个子串，n表示替换的次数（-1表示全部替换）。 
6. strings.Count(str, substr) int：返回子串在字符串中出现的次数。

## 读取文件

逐行读取
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	// 延迟关闭文件
	defer file.Close()

	// 使用 bufio 包读取文件
	scanner := bufio.NewScanner(file)

	// 逐行读取文件
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
```
### 读取并解析CSV和TSV
对于CSV和TSV文件，可以使用Go语言内置的encoding/csv包来进行读取和解析。该包提供了一个Reader类型，它可以逐行读取CSV和TSV文件，并返回一个二维数组，每一行对应一个子数组。

```go
package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    // 打开 CSV 文件
    f, err := os.Open("data.csv")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // 创建 CSV 读取器
    r := csv.NewReader(f)
    r.Comma = ',' // 设置分隔符为逗号，对于TSV文件，只需要将r.Comma设置为制表符即可
// r.Comma = '\t' // 设置分隔符为制表符
    // 逐行读取 CSV 文件
    for {
        record, err := r.Read()
        if err != nil {
            break
        }
        fmt.Println(record)
    }
}
值得注意的是，r.Read方法返回的是一个由字符串组成的切片，每个字符串都是CSV或TSV文件中的一个字段。如果需要将字符串转换为其他类型，例如数字或布尔值，可以使用Go语言内置的转换函数进行转换。
```

## 写入文件
### 逐行写入
该代码会将10行数据写入名为test.txt的文件中。其中，os.OpenFile()函数用于打开文件，os.O_CREATE|os.O_WRONLY|os.O_TRUNC参数表示如果文件不存在则创建，写入方式为覆盖写入，文件权限为0644。创建写入器时使用了bufio.NewWriter()函数，该函数接受一个文件操作对象作为参数，返回一个写入器对象。使用写入器对象的WriteString()方法逐行写入数据，写入完毕后使用Flush()方法将缓存中的数据写入文件。
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开文件，如果不存在则创建
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Open file error:", err)
		return
	}
	defer file.Close()

	// 创建写入器
	writer := bufio.NewWriter(file)

	// 写入数据
	for i := 1; i <= 10; i++ {
		_, err := writer.WriteString(fmt.Sprintf("line %d\n", i))
		if err != nil {
			fmt.Println("Write file error:", err)
			return
		}
	}

	// 将缓存中的数据写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("Flush buffer error:", err)
		return
	}
}
```

### 逐行写入TSV文件
1. 使用bufio.NewWriter创建一个带缓冲的写入器writer，可以避免频繁地写入磁盘，提高写入效率。 
2. 使用strings.Join函数将每一行的数据用\t连接起来，形成一行tsv格式的数据。 
3. 在写入每一行数据后，通过writer.Flush()强制将数据写入磁盘。
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    data := [][]string{
        {"Name", "Age", "Gender"},
        {"Tom", "20", "Male"},
        {"Alice", "25", "Female"},
        {"Bob", "30", "Male"},
    }

    file, err := os.Create("output.tsv")
    if err != nil {
        fmt.Println("Failed to create file:", err)
        return
    }
    defer file.Close()

    writer := bufio.NewWriter(file)

    for _, row := range data {
        line := strings.Join(row, "\t")
        _, err := writer.WriteString(line + "\n")
        if err != nil {
            fmt.Println("Failed to write data:", err)
            return
        }
    }

    writer.Flush()
    fmt.Println("Data has been written to the file.")
}
```

### 逐行写入csv文件
```go
package main

import (
	"bufio"
	"encoding/csv"
	"os"
)

func main() {
	data := [][]string{
		{"first_name", "last_name", "email"},
		{"John", "Doe", "john@example.com"},
		{"Jane", "Doe", "jane@example.com"},
		{"Joe", "Schmoe", "joe@example.com"},
	}

	file, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(bufio.NewWriter(file))
	defer writer.Flush()

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
}
```






