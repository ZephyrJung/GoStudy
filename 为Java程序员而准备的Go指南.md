# 为Java程序员而准备的Go指南

> https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/

这篇文章是为了帮助Java程序员们迅速的掌握Go语言。

本篇将先用Java程序员耳熟能详的特性举例，then gives a fairly detailed description of Go’s building blocks,and ends with an example illustrating constructsthat have no immediate counterpart in Java.

## Hello Stack (example)

为了激发你的兴趣，我们将以一个麻雀虽小五脏俱全的典型案例开始，即Stack.java

Go语言实现如下：

```java
// Package collection implements a generic stack.
package collection
// The zero value for Stack is an empty stack ready to use.
type Stack struct {
    data []interface{}
}
// Push adds x to the top of the stack.
func (s *Stack) Push(x interface{}) {
    s.data = append(s.data, x)
}
// Pop removes and returns the top element of the stack.
// It’s a run-time error to call Pop on an empty stack.
func (s *Stack) Pop() interface{} {
    i := len(s.data) - 1
    res := s.data[i]
    s.data[i] = nil  // to avoid memory leak
    s.data = s.data[:i]
    return res
}
// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
    return len(s.data)
}
```

- 在最顶部声明上的注释是文档注释，用纯文本书写
- 声明的名字写在type之后
- `struct`类似于Java中的class，但struct中的成员不能是方法，只可以是变量
- `interface{}`类似于Java中的`Object`。不过它被所有的Go类型实现，不仅仅是引用类型。
- 代码段`(s *Stack)`声明了一个方法调用者`s`，类似于Java中的`this`
- 操作符`:=`同时声明并初始化变量，类型通过初始化表达式的值进行推断

如下是一个Hello World程序，展示如何使用`collection.Stack`这个抽象数据类型。

```go
package collection_test
import (
    collection "."
    "fmt"
)
func Example() {
    var s collection.Stack
    s.Push("world")
    s.Push("hello, ")
    for s.Size() > 0 {
        fmt.Print(s.Pop())
    }
    fmt.Println()
    // Output: hello, world
}
```

这个`cllection_test`测试包与`collection`包的位置在同一目录下。第一个import声明意味着我们将使用当前目录(`"."`)下的包，并赋予一个别名`collection`。第二个声明包含指向标准包路径(`"fmt"`)；如果没有指定别名，则实际包名`fmt`将作为默认名称。

## 概念上的区别

- Go没有带构造器的类，实例方法，继承层次结构，动态方法查找，取而代之的是structs和interfaces。Interfaces也可以用在类似Java的泛型的地方。
- Go提供任何类型的值的指针，不仅仅是对象和数组。对于任何类型`T`，都有一个对应的指针类型`*T`，表示该类型值的指针。
- Go允许任意类型上的方法，无需装箱。方法的调用者（类似于Java中的this），可以是一个值或一个指针。
- 在Go中，数组属于数值。当数组被用作方法参数时，方法接收到的是一个数组的拷贝，而非指针。然而实践中，方法经常使用slices作为参数。slices是对数组的底层引用。
- 语言内置支持字符串， a string behaves like a slice of bytes, 但它是不可改变的。
- 语言内置支持哈希表，称为map
- 单独执行线程，goroutines，以及线程之间的通信渠道，channels，都被Go语言内置支持。
- 特定类型（maps，slices，channels）是通过引用而非值传递的。这就是说，想方法传递一个map类型并不是传递了map的拷贝。如果方法改变了这个map，方法外部调用者也会看到。用Java来讲，可以把这个想象成map的引用。
- Go提供了两种访问级别，与Java的public和package private类似。名称首部为大写时代表public，否则为package-private
- 并没有使用exceptions，Go使用了error类型来代表注入到达文件末尾之类的事件，以及运行时panics来代表注入试图越界读取数组的运行时错误
- Go不支持隐式类型转换。包含不同类型的表达式需要进行显式转换
- Go不支持方法重载。函数和方法在同一作用域内必须有不同的名称
- Go用`nil`代表错误的指针，类似于Java的`null`

## 语法

### 声明

声明语法与Java正好相反，将类型写在名称的后面。类型声明从左往右读可能更容易点儿。

| Go                           | Approximate Java equivalent              |
| ---------------------------- | ---------------------------------------- |
| `var v1 int`                 | `int v1 = 0;`                            |
| `var v2 *int`                | `Integer v2 = null;`                     |
| `var v3 string`              | `String v3 = "";`                        |
| `var v4 [10]int`             | `int[] v4 = new int[10]; // v4 is a value in Go.` |
| `var v5 []int`               | `int[] v5 = null;`                       |
| `var v6 *struct { a int }  ` | `C v6 = null; // Given: class C { int a; }` |
| `var v7 map[string]int`      | `HashMap v7 = null;`                     |
| `var v8 func(a int) int`     | `F v8 = null; // interface F { int f(int a); }` |

声明通常是以关键字后面跟随着被声明的对象名称的形式，关键字可能是`const`，`type`，`var`，或`func`。你可以用一个关键字把一系列的声明写在括号内。

```go
var(
  n int
  x float64
)
```

当声明一个方法时，你必须要么为所有参数提供名称，要么一个名称也不写，不能忽略了一些命名，而又提供了另一部分的名称。可以对相同类型的名称进行分组：

```go
func f(i,j,k int,s,t string)
```

变量可以在声明的时候初始化。如果初始化了，仍然可以进行类型定义但是没有必要，它将默认位初始化表达式的值类型。

```go
var v9=*v2
```

如果变量没有被显式初始化，那么必须要指定类型。防止它被隐式初始化为零值（`0`，`nil`，`""`，等）。Go里没有未初始化的变量。

### 简单声明

在方法内部，可以使用简化的声明语法`:=`。

`v10:=v1`等价于`var v10=v1`

### 方法类型

在Go中，方法是一等公民。Go的方法类型表示一系列相同参数和返回结果类型的方法。

```go
type binOp func(int,int) int
var op binOp
add :=func(i,j int) int{return i+j}
op=add
n=op(100,200)
```

### 多重赋值

Go允许多重复值，右边的表达式将在对操作符左边进行赋值前求值。

```go
i,j=j,i //替换i和j
```

方法可能有多个返回值，由括号内的列表指出。返回值可以赋给变量列表。

```go
func f() (i int,pj *int){...}
v1,v2=f()
```

### 空白标识

