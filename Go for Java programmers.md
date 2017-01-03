# Go for Java programmers

> https://www.nada.kth.se/~snilsson/go_for_java_programmers/

This text is intended to help Java programmerscome up to speed quickly with [Go](http://golang.org).

It starts with an example highlighting featureseasily recognized by all Java programmers,then gives a fairly detailed description of Go’s building blocks,and ends with an example illustrating constructsthat have no immediate counterpart in Java.

## Hello stack (example)

To whet your appetite, we start with a small but complete and idiomatic examplecorresponding to this [Stack.java](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/src/collection/Stack.java) program.

```
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

[stack.go](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/src/collection/stack.go)

- Comments that appear directly before top-level declarations are    documentation comments. They are written in plain text.
- For declarations, your write the name followed by the type.
- A `struct` corresponds to a class in Java,    but the members of a struct cannot be methods, only variables.
- The type `interface{}` corresponds to Java’s    `Object`. However, it is implemented by all types in Go,    not only reference types.
- The code fragment `(s *Stack)`    declares a method receiver `s`    corresponding to Java’s `this`.
- The operator `:=` both declares and initializes a variable.    Its type is deduced from the initialization expression.

And here is a Hello world programdemonstrating how to use the `collection.Stack` abstract data type.

```
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

[example_test.go](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/src/collection/example_test.go)

The test package `collection_test` resides in the same directoryas the `collection` package.The first import declaration indicates that we will use the package from the currentdirectory (`"."`) and that we will refer to it bythe name `collection` within this file.The second import declaration contains the path (`"fmt"`) to astandard package; since no name is given, the actual package name,`fmt`, will be used by default.

## Conceptual differences

- Go does not have classes with constructors.    Instead of instance methods, a class inheritance hierarchy,    and dynamic method lookup, Go provides    [structs](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Structs) and    [interfaces](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Methods_and_interfaces).    Interfaces are also used where Java uses generics.
- Go offers pointers to values of all types, not just objects and arrays.    For any type `T`, there is a corresponding pointer type    `*T`, denoting [pointers](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Pointers)    to values of type `T`.
- Go allows methods on any type; no boxing is required.     The method *receiver*,     which corresponds to `this` in Java,     can be a direct value or a pointer.
- Arrays in Go are values. When an array is used as a    function parameter, the function receives a copy of the array,    not a pointer to it. However, in practice functions often use     [slices](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Slices) for parameters;    slices are references to underlying arrays.
- Strings are provided by the language; a string behaves like a     slice of bytes, but is immutable.
- Hash tables are provided by the language. They are called maps.
- Separate threads of execution, [goroutines](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Goroutines),    and communication channels between them, [channels](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Goroutines),    are provided by the language.
- Certain types (maps, slices, and channels) are passed by reference,    not by value. That is, passing a map to a function does not copy the map;    if the function changes the map, the change will be seen by the caller.    In Java terms, one can think of this as being a reference to the map.
- Go provides two access levels, analogous to Java’s public and    package-private. Top-level declarations are public if their names    start with an upper-case letter, otherwise they are package-private.
- Instead of exceptions, Go uses values of type    [error](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Errors) to signify events such as end-of-file,    and run-time [panics](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Panic) for run-time errors    such as attempting to index an array out of bounds.
- Go does not support implicit type conversion. Operations that mix    different types require an explicit conversion.
- Go does not support function overloading.     Functions and methods in the same scope must have unique names.
- Go uses `nil` for invalid pointers, where Java uses    `null`.

## Syntax

  [    ![Syntax Terror](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/syntax-terror.jpg)  ](http://www.flickr.com/photos/ruminatrix/3052493260/)

### Declarations

The declaration syntax is reversed compared to Java.You write the name followed by the type.Type declarations may be read easily from left to right.

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

Declarations generally take the form of a keyword followed by the nameof the object being declared. The keyword is one of `const`,`type`, `var`, or `func`.You can also use a keyword followed by a series of declarations inparentheses.

```
var (
    n int
    x float64
)

```

When declaring a function, you must either provide a name for each parameteror not provide a name for any parameter; you can’t omit some namesand provide others.  You may group several names with the same type:

```
func f(i, j, k int, s, t string)

```

A variable may be initialized when it is declared.  When this is done,specifying the type of the variable is permitted but not required.When the type is not specified, it defaults to the type ofthe initialization expression.

```
var v9 = *v2

```

If a variable is not initialized explicitly, the type must be specified.In that case it will be implicitly initialized to the type’s*zero value*(`0`, `nil`, `""`, etc.).There are no uninitialized variables in Go.

### Short declarations

Within a function, a short declaration syntax is availablewith `:=` .

```
v10 := v1

```

This is equivalent to

```
var v10 = v1

```

### Function types

In Go, functions are first-class citizens.Go’s function type denotes the set of all functionswith the same parameter and result types.

```
type binOp func(int, int) int

var op binOp
add := func(i, j int) int { return i + j }

op = add
n = op(100, 200)  // n = 100 + 200

```

### Multiple assignment

Go permits multiple assignments.The expressions on the right are evaluatedbefore assigning to any of the operands on the left.

```
i, j = j, i  // Swap i and j.

```

Functions may have multiple return values, indicated by a list inparentheses.  The returned values can be stored by assignmentto a list of variables.

```
func f() (i int, pj *int) { ... }
v1, v2 = f()

```

### The blank identifier

The blank identifier, represented by the underscore character,provides a way to ignore values returned by a multi-valued expression:

```
v1, _ = f()  // Ignore second value returned by f().

```

### Semicolons and formatting

Instead of worrying about semicolons and formatting, you may usethe `gofmt` program to produce a single standard Go style.While this style may initially seem odd, it is as good as any other style,and familiarity will lead to comfort.

Go code uses very few semicolons in practice. Technically, all Gostatements are terminated by a semicolon. However, Goimplicitly inserts a semicolon at the end of a non-blank lineunless the line is clearly incomplete.A consequence of this is that in some cases Go does not permit a line break.For example, you may not write

```
func g()
{            // INVALID; "{" should be on previous line.
}

```

A semicolon will be inserted after `g()`causing it to be a function declaration rather than a function definition.Similarly, you may not write

```
if n == 0 {
}
else {       // INVALID; "else {" should be on previous line.
}

```

A semicolon will be inserted after the `}` precedingthe `else`, causing a syntax error.

### Conditional statements

Go does not use parentheses around the condition ofan `if` statement,or the expressions of a `for` statement,or the value of a `switch` statement.On the other hand, it does require curly braces aroundthe body of an `if` or `for` statement.

```
if a < b { f() }
if (a < b) { f() }           // Parentheses are unnecessary.
if (a < b) f()               // INVALID
for i = 0; i < 10; i++ {}
for (i = 0; i < 10; i++) {}  // INVALID

```

Furthermore, `if` and `switch` acceptan optional initialization statement, which is commonly usedto set up a local variable.

```
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}

```

### For statements

Go does not have a `while` statementnor does it have a `do-while` statement.The `for` statement may be used with a single condition,which makes it equivalent to a `while` statement.Omitting the condition entirely produces an endless loop.

A `for` statement may contain a `range` clausefor iterating over strings, arrays, slices, maps, or channels.Instead of writing

```
for i := 0; i < len(a); i++ { ... }

```

to loop over the elements of `a`, we could also write

```
for i, v := range a { ... }

```

This assigns `i` to the index and `v`to the value of the successive elements of an array, slice, or string.For strings, `i` is an index to a byte and`v` is a Unicode code point of type `rune`(`rune` is an alias for `int32`).Iterations over maps produce key-value pairs,while channels produce only one iteration value.

### Break and continue

Like Java, Go permits `break` and `continue`to specify a label, but the label must refer to a `for`,`switch`, or `select` statement.

### Switch statements

In a `switch` statement, `case` labelsdo not fall through by default,but you can make them fall through by ending a casewith a `fallthrough` statement.

```
switch n {
case 0:  // empty case body
case 1:
    f()  // f is not called when n == 0.
}

```

But a `case` can have multiple values.

```
switch n {
case 0, 1:
    f()  // f is called if n == 0 || n == 1.
}

```

The values in a `case` can be any type that supportsthe equality comparison operator, such as strings or pointers.A missing switch expression is equivalent tothe expression `true`.

```
switch {
case n < 0:
    f1()
case n == 0:
    f2()
default:
    f3()
}

```

### ++ and -- statements

The `++` and `--` may only be usedas postfix operators and only in statements, not in expressions.For example, you cannot write `n = i++`.

### The defer statement

A `defer` statement invokes a function whose execution isdeferred to the moment the surrounding function returns.The deferred function will be executed regardless of which paththe surronding function takes to return.The parameters of the deferred function, however, are computed and savedfor future use already when the `defer` statement executes.

```
f, err := os.Open("filename")
defer f.Close()  // f will be closed when this function returns.

```

## Constants

In Go constants may be *untyped*.This applies to numeric literals, expressions using only untyped constants,and `const` declarations where no type is given andthe initializer expression is untyped.A value derived from an untyped constant becomes typed when itis used within a context that requires a typed value.This permits constants to be used relatively freely even thoughGo has no implicit type conversion.

```
var a uint
f(a + 1)    // The untyped numeric constant 1 becomes typed as uint.
f(a + 1e3)  // 1e3 is also typed as uint.

```

The language does not impose any limits on the size of an untypednumeric constant. A limit is only applied when a constant is usedwhere a type is required.

```
const huge = 1 << 100
var n int = huge >> 98

```

If the type is absent in a variable declaration and the correspondingexpression evaluates to an untyped numeric constant,the constant is converted to type `rune`, `int`,`float64`, or `complex128` respectively,depending on whether the value is a character, integer, floating-point, orcomplex constant.

```
c := 'å'      // rune (alias for int32)
n := 1 + 2    // int
x := 2.7      // float64
z := 1 + 2i   // complex128

```

Go does not have enumerated types.Instead, you can use the special name `iota` in a single`const` declaration to get a series of increasing value.When an initialization expression is omitted for a `const`,it reuses the preceding expression.

```
const (
    red = iota  // red == 0
    blue        // blue == 1
    green       // green == 2
)

```

## Structs

A struct corresponds to a class in Java, but the members of a structcannot be methods, only variables.A pointer to a struct is similar to a reference variable in Java.As opposed to Java classes, structs may also be defined as direct values.In both cases you use `.` to access the members of a struct.

```
type MyStruct struct {
    s string
    n int64
}

var x MyStruct      // x is initialized to MyStruct{"", 0}.
var px *MyStruct    // px is initialized to nil.
px = new(MyStruct)  // px points to the new struct MyStruct{"", 0}.

x.s = "Foo"
px.s = "Bar"

```

In Go, methods may be associated with any named type, not just with structs;see the [discussion of methods and interfaces](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Methods_and_interfaces).

## Pointers

  [    ![Pointers](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/pointers.jpg)  ](http://www.flickr.com/photos/lwr/374707207/)

If you have an int or a struct or an array,assignment copies the contents of the object.To achieve the effect of Java reference variables, Go uses pointers.For any type `T`, there is a correspondingpointer type `*T`,denoting pointers to values of type `T`.

To allocate storage for a pointer variable,use the built-in function `new`,which takes a type and returns a pointer to the allocated storage.The allocated space will be zero-initialized for the type.For example, `new(int)` allocates storagefor a new `int`,initializes it with the value `0`,and returns its address, which has type `*int`.

The Java code `T p = new T()`,where `T` is a class with two instance variables`a` and `b` of type `int`,corresponds to

```
type T struct { a, b int }
var p *T = new(T)

```

or the more idiomatic

```
p := new(T)

```

The declaration `var v T`,which declares a variable that holds a value of type `T`,has no equivalent in Java.Values can also be created and initialized using a *composite literal*.For example:

```
v := T{1, 2}

```

is equivalent to

```
var v T
v.a = 1
v.b = 2

```

For an operand `x` of type `T`,the *address operator* `&x`gives the address of `x`, a value of type `*T`.For example:

```
p := &T{1, 2} // p has type *T

```

For an operand `x` of pointer type,the *pointer indirection* `*x`denotes the value pointed to by `x`.Pointer indirections are rarely used;Go, just like Java, can automatically take the address of a variable:

```
p := new(T)
p.a = 1 // equivalent to (*p).a = 1

```

## Slices

A slice is conceptually a struct with three fields:a pointer into an array, a length, and a capacity.Slices support the `[]` operatorto access elements of the underlying array.The built-in `len` function returns the length of the slice.The built-in `cap` function returns the capacity.

Given an array, or another slice, a new slice is created via`a[i:j]`. This creates a new slice which refers to `a`,starts at index `i`, and ends before index `j`.It has length `j - i`.If `i` is omitted, the slice starts at `0`.If `j` is omitted, the slice ends at `len(a)`.The new slice refers to the same array to which `a` refers.That is, changes made to the elements using the new slicemay be seen using `a`.The capacity of the new slice is simply the capacity of `a`minus `i`.The capacity of an array is the length of the array.

```
var s []int
var a [10]int

s = a[:]  // short for s = a[0:len(a)]

```

If you create a value of type `[100]byte`(an array of 100 bytes, perhaps a buffer)and you want to pass it to a function without copying it,declare the function parameter to have type `[]byte`,and pass a slice of the array.Slices may also be created using the `make` functionas [described below](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Making_values).

Slices combined with the built-in function `append` offermuch the same functionality as Java’s `ArrayList`.

```
s0 := []int{1, 2}
s1 := append(s0, 3)      // append a single element
s2 := append(s1, 4, 5)   // append multiple elements
s3 := append(s2, s0...)  // append a slice

```

The slice syntax may also be used with a string. It returns a new stringwhose value is a substring of the original string.

## Making values

Map and channel values must be allocated using the built-in function`make`. For example, calling

```
make(map[string]int)

```

returns a newly allocated value of type `map[string]int`.As opposed to `new`, `make` returnsthe actual object, not an address. This is consistent with the factthat maps and channels are reference types.

For maps, make takes a capacity hint as an optional second argument.For channels, there is an optional second argument that sets thebuffering capacity of the channel; the default is 0 (unbuffered).

The `make` function may also be used to allocate a slice.In this case it allocates memory for the underlying array and returnsa slice referring to it.There is one required argument, which is the number of elements in the slice.A second optional argument is the capacity of the slice.

```
m := make([]int, 10, 20)  // Same as new([20]int)[:10]

```

## Methods and interfaces

### Methods

A method looks like an ordinary function definition,except that it has a *receiver*.The receiver is similar to the `this` referencein a Java instance method.

```
type MyType struct { i int }

func (p *MyType) Get() int {
    return p.i
}

var pm = new(MyType)
var n = pm.Get()

```

This declares a method `Get` associated with `MyType`.The receiver is named `p` in the body of the function.

Methods are defined on *named types*. If you convert the valueto a different type, the new value will have the methods of the new type,not those of the old type.

You may define methods on a built-in type by declaring a new named typederived from it. The new type is distinct from the built-in type.

```
type MyInt int

func (p MyInt) Get() int {
    return int(p)  // The conversion is required.
}

func f(i int) {}
var v MyInt

v = v * v          // The operators of the underlying type still apply.
f(int(v))          // int(v) has no defined methods.
f(v)               // INVALID

```

### Interfaces

A Go interface is similar to a Java interface,but any type that provides the methods named in a Go interfacemay be treated as an implementation of that interface. No explicit declaration is required.

Given this interface:

```
type MyInterface interface {
    Get() int
    Set(i int)
}

```

Since `MyType` already has a `Get` method,we can make `MyType` satisfy the interface by adding

```
func (p *MyType) Set(i int) {
    p.i = i
}

```

Now any function which takes `MyInterface` as a parameterwill accept a variable of type `*MyType`.

```
func GetAndSet(x MyInterface) {}

func f1() {
    var p MyType
    GetAndSet(&p)
}

```

In Java terms, defining `Set` and `Get` for`*MyType` made `*MyType` automatically implement`MyInterface`. A type may satisfy multiple interfaces.This is a form of duck typing.

> When I see a bird that walks like a duck and swimslike a duck and quacks like a duck, I call that bird a duck.

James Whitcomb Riley

### Anonymous fields

An anonymous field may be used to implement something much like a Java subclass.

```
type MySubType struct {
    MyType
    j int
}

func (p *MySubType) Get() int {
    p.j++
    return p.MyType.Get()
}

```

This effectively implements `MySubType` as a subtype of`MyType`.

```
func f2() {
    var p MySubType
    GetAndSet(&p)
}

```

The `Set` method is inherited from `MyType`,because methods associated with the anonymous field are promotedto become methods of the enclosing type.In this case, because `MySubType` has an anonymous fieldof type `MyType`, the methods of `MyType` alsobecome methods of `MySubType`.The `Get` method was overridden,and the `Set` method was inherited.

This is not precisely the same as a subclass in Java.When a method of an anonymous field is called, its receiver is the field,not the surrounding struct. In other words, methods on anonymous fieldsare not dynamically dispatched. When you want the equivalent of Java’sdynamic method lookup, use an interface.

```
func f3() {
    var v MyInterface

    v = new(MyType)
    v.Get()  // Call the Get method for *MyType.

    v = new(MySubType)
    v.Get()  // Call the Get method for *MySubType.
}

```

### Type assertions

A variable that has an interface type may be converted to have adifferent interface type using a type assertion.This is implemented dynamically at run time. Unlike Java, there does notneed to be any declared relationship between the two interfaces.

```
type Printer interface {
    Print()
}

func f4(x MyInterface) {
    x.(Printer).Print()  // type assertion to Printer
}

```

The conversion to `Printer` is entirely dynamic.It will work as long as the *dynamic type* of x(the actual type of the value stored in `x`)defines a `Print` method.

### Generics

Go doesn’t have generic types, but by combining anonymous fieldsand type assertions it’s possible to achieve something akin toJava’s parameterized types.

```
type StringStack struct {
    Stack
}

func (s *StringStack) Push(n string) { s.Stack.Push(n) }
func (s *StringStack) Pop() string   { return s.Stack.Pop().(string) }

```

`StringStack` specializes the generic `Stack`from the [Hello stack](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Hello) exampleso that it operates on `string` elements only ‒just like `Stack` in Java.Notice that the `Size` method is inheritedfrom `Stack`.

## Errors

Where Java typically uses exceptions, Go has two different mechanisms.Most functions return *errors*; only truly unrecovorable conditions,such as an out-of-range index, produce run-time exceptions.

Go’s multivalued return makes it easy to return a detailed error messagealongside the normal return value. By convention, such messageshave type `error`, a simple built-in interface.

```
type error interface {
    Error() string
}

```

For example, the `os.Open` function returns a non-nil`error` value when it fails to open a file.

```
func Open(name string) (file *File, err error)

```

The following code uses `os.Open` to open a file.If an `error` occurs it calls `log.Fatal`to print the error message and stop.

```
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f

```

The `error` interface requires only an `Error` method,but specific `error` implementations often have additional methods,allowing callers to inspect the details of the error.

## Panic and recover

A *panic* is a run-time errorthat unwinds the stack of the goroutine,running any deferred functions along the way,and then stops the program.Panics are similar to Java exceptions,but are only intended for run-time errors,such as following a `nil` pointer orattempting to index an array out of bounds.To signify events such as end-of-file, Go programs use the built-in`error` type as [described above](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/#Errors).

The built-in function `recover` can be used to regaincontrol of a panicking goroutine and resume normal execution.A call to `recover` stops the unwinding and returnsthe argument passed to `panic`.Because the only code that runs while unwinding is inside deferred functions,`recover` is only useful inside deferred functions.If the goroutine is not panicking,`recover` returns `nil`.

## Goroutines and channels

  [    ![Sushi conveyor belt](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/sushi-conveyor-belt.jpg)  ](http://www.flickr.com/photos/erikjaeger/35008017/)        Sushi conveyor belt  

### Goroutines

Go permits starting a new thread of execution, a goroutine,using the `go` statement.It runs a function in a different, newly created, goroutine.All goroutines in a single program share the same address space.

Goroutines are lightweight,costing little more than the allocation of stack space.The stacks start small and grow by allocating and freeingheap storage as required.Internally goroutines act like coroutines that are multiplexedamong multiple operating system threads.You do not have to worry about these details.

```
go list.Sort()  // Run list.Sort in parallel; don’t wait for it. 

```

Go has function literals, which can act as *closures*and are powerful when coupled with the `go` statement.

```
// Publish prints text to stdout after the given time has expired.
func Publish(text string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(text)
    }()  // Note the parentheses. We must call the function.
}

```

The variables `text` and `delay`are shared between the surrounding function and the function literal;they survive as long as they are accessible.

### Channels

A channel provides a mechanism for two goroutines to synchronizeexecution and communicate by passing a value of a specifiedelement type.The `<-` operator specifies the channel direction,send or receive. If no direction is given, the channel is bi-directional.

```
chan Sushi      // can be used to send and receive values of type Sushi
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints

```

Channels are a reference type and are allocated with make.

```
ic := make(chan int)        // unbuffered channel of ints
wc := make(chan *Work, 10)  // buffered channel of pointers to Work

```

To send a value on a channel,use `<-` as a binary operator.To receive a value on a channel, use it as a unary operator.

```
ic <- 3       // Send 3 on the channel.
work := <-wc  // Receive a pointer to Work from the channel.

```

If the channel is unbuffered,the sender blocks until the receiver has received the value.If the channel has a buffer,the sender blocks only until the value has been copied to the buffer;if the buffer is full,this means waiting until some receiver has retrieved a value.Receivers block until there is data to receive.

The `close` function records that no more valueswill be sent on a channel. After calling `close`,and after any previously sent values have been received,receive operations will return a zero value without blocking.A multi-valued receive operation additionally returnsan indication of whether the channel is closed.

```
ch := make(chan string)
go func() {
    ch <- "Hello!"
    close(ch)
}()
fmt.Println(<-ch)  // Print "Hello!".
fmt.Println(<-ch)  // Print the zero value "" without blocking.
fmt.Println(<-ch)  // Once again print "".
v, ok := <-ch      // v is "", ok is false.

```

In the next example we let the `Publish` functionreturn a channel, which is used to broadcast a message whenthe text has been published.

```
// Publish prints text to stdout after the given time has expired.
// It closes the wait channel when the text has been published.
func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
    ch := make(chan struct{})
    go func() {
        time.Sleep(delay)
        fmt.Println(text)
        close(ch)
    }()
    return ch
}

```

This is how you might use this `Publish` function.

```
wait := Publish("important news", 2 * time.Minute)
// Do some more work.
<-wait // blocks until the text has been published

```

### Select statement

The select statement is the final tool in Go’s concurrency toolkit.It chooses which of a set of possible communications will proceed.If any of the communications can proceed, one of them is randomlychosen and the corresponding statements are executed.Otherwise, if there is no default case,the statement blocks until one of the communications can complete.

Here is a toy example showing how the select statement canbe used to implement a random number generator.

```
rand := make(chan int)
for { // Send random sequence of bits to rand.
    select {
    case rand <- 0: // note: no statement
    case rand <- 1:
    }
}

```

Somewhat more realistically, here is how a select statementcould be used to set a time limit on a receive operation.

```
select {
case news := <-AFP:
    fmt.Println(news)
case <-time.After(time.Minute):
    fmt.Println("Time out: no news in one minute.")
}

```

The function `time.After` is part of the standard library;it waits for a specified time to elapse and then sends the current timeon the returned channel.

## Concurrency (example)

We end with a small but complete example to show how the pieces fit together.It’s draft code for a server accepting `Work` requeststhrough a channel. Each request is served in a separate goroutine.The `Work` struct itself contains a channelused to return a result.

```
package server

import "log"

// New creates a new server that accepts Work requests
// through the req channel.
func New() (req chan<- *Work) {
    wc := make(chan *Work)
    go serve(wc)
    return wc
}

type Work struct {
    Op    func(int, int) int
    A, B  int
    Reply chan int  // Server sends result on this channel.
}

func serve(wc <-chan *Work) {
    for w := range wc {
        go safelyDo(w)
    }
}

func safelyDo(w *Work) {
    // Regain control of panicking goroutine to avoid
    // killing the other executing goroutines.
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(w)
}

func do(w *Work) {
    w.Reply <- w.Op(w.A, w.B)
}

```

[server.go](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/src/server/server.go)

And this is how you might use it.

```
package server_test

import (
    server "."
    "fmt"
    "time"
)

func main() {
    s := server.New()

    divideByZero := &server.Work{
        Op:    func(a, b int) int { return a / b },
        A:     100,
        B:     0,
        Reply: make(chan int),
    }
    s <- divideByZero

    select {
    case res := <-divideByZero.Reply:
        fmt.Println(res)
    case <-time.After(time.Second):
        fmt.Println("No result in one second.")
    }
    // Output: No result in one second.
}

```

[example_test.go](https://www.nada.kth.se/%7Esnilsson/go_for_java_programmers/src/server/example_test.go)

Concurrent programming is a large topic andGo’s approach is quite different from Java’s.Here are two texts that cover the basics.

- [Fundamentals of concurrent programming](https://www.nada.kth.se/%7Esnilsson/concurrency/)is a gentle introduction to concurrency with small examples in Go.
- [Share Memory by Communicating](http://golang.org/doc/codewalk/sharemem/)is a codewalk with a more substantial example.