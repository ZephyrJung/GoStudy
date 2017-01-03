# 为Java程序员而准备的Go指南

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

如下是一个Hello World程序，展示如何使用`cllection.Stack`这个抽象数据类型。

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

