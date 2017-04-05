# Go by Example

## Hello World

我们第一个程序就是打印经典的“hello world”，下面是完整的代码

```go
package main
import "fmt"
func main(){
  fmt.Println("hello world")
}
```

要运行这个程序，将代码保存为hello-world.go，然后使用`go run`

有时候我们想让程序编译成二进制文件，可以使用`go build`，然后就可以直接运行了。

现在我们已经运行及编译了基本的Go程序，让我们了解更多关于这个语言的东西吧。

## Values

Go有多种值的类型，包括string，integer，float，boolean等。如下是几个基本例子。

```go
package main
import "fmt"
func main(){
  //string可以使用+连接在一起
  fmt.Println("go"+"lang")
  fmt.Println("1+1=",1+1)
  fmt.Pritnln("7.0/3.0=",7.0/3.0)
  fmt.Println(true&&false)
  fmt.Println(true||false)
  fmt.Pritnln(!true)
}
```

## Variables

在Go中，变量被编译器显式的声明和使用，例如检查函数调用类型的正确性。

```go
package main
import "fmt"
func main(){
  //声明一个或多个变量
  var a string="initial"
  fmt.Println(a)
  var b,c int = 1,2
  fmt.Println(b,c)
  //Go将推断变量的初始化类型
  var d = true
  fmt.Println(d)
  //声明没有初始值的变量将被初始化为零值，如int的零值是0
  var e int
  fmt.Println(e)
  //var f string = "short"的简化写法
  f:="short"
  fmt.Println(f)
}
```

## Constants

Go支持的常量有字符、字符串、布尔值以及数值

```go
package main
import "fmt"
import "math"
const s string = "constant"
func main(){
  fmt.Println(s)
  //const声明可以出现在任何var声明出现的地方
  const n = 50000000
  //常量表达式可以任意精度进行运算
  const d = 3e20 / n
  fmt.Println(d)
  //数值常量没有类型，除非进行了显式转换等类型赋予
  fmt.Prinln(int64(d))
  //数值在使用环境上下文中会得到类型，如变量赋值或者方法调用，如math.Sin需要的是一个float64
  fmt.Println(math.Sin(n))
}
```

## For

for是Go中唯一的循环结构，下面是三种基本的for循环类型。

```go
package main
import "fmt"
func main(){
  i := 1
  //最基本的类型，只有一个条件
  for i<=3{
    fmt.Println(i)
    i = i+1
  }
  //经典的 初始化/条件/循环后 for循环
  for j:=7;j<=9;j++{
    fmt.Println(j)
  }
  //没有条件的for语句将无限循环，除非在内部的方法中使用了break或者return
  for{
    fmt.Println("loop")
    break
  }
  //也可以使用continue来直接到下一个循环
  for n:=0;n<=5;n++ {
    if n%2 == 0{
      continue
    }
    fmt.Println(n)
  }
}
```

## If/Else

Go中的if-else分支简单直接。

```go
package main
import "fmt"
func main(){
  if 7%2==0{
    fmt.Println("7 is even")
  }else{
    fmt.Println("7 is odd")
  }
  //可以单独使用if
  if 8%4==0{
    fmt.Println("8 is divisible by 4")
  }
  //条件之前可以有语句，声明在该语句中的任何变量可以用于所有分支
  if num:=9;num<0{
    fmt.Println(num,"is negative")
  }else if num<10{
    fmt.Println(num,"has 1 digit")
  }else{
    fmt.Println(num,"has multiple digits")
  }
}
```

注意，Go中条件周围不需要圆括号，但是花括号是必要的。

Go中没有三目if语句，所以对于最基本的条件也需要写完整的if语句。

## Switch

```go
package main
import "fmt"
import "time"
func main(){
  i:=2
  fmt.Print("write",i," as ")
  switch i{
  case 1: fmt.Println("one")
  case 2: fmt.Println("two")
  case 3: fmt.Println("three")
  }
  //可以在同一个case语句中使用逗号来分隔多个表达式
  switch time.Now().weekday(){
  case time.Saturday,time.Sunday:
    fmt.Println("it's the weekend")
  //这里也使用了default case
  default:
    fmt.Println("it's a weekday")
  }
  t:=time.Now()
  //没有表达式的switch是另一种实现if/else逻辑的路子，同时这也展示了case表达式可以是非常量值。
  switch{
  case t.Hour()<12:
    fmt.Println("it's before noon")
  default:
    fmt.Println("it's after noon")
  }
  //类型switch比较了类型而非值
  whatAmI := func(i interface{}){
    switch t:=i.(type){
    case bool:
      fmt.Println("I'm a bool")
    case int:
      fmt.Println("I'm an int")
    default:
      fmt.Println("Don't know type %T\n",t)
    }
  }
  whatAmI(true)
  whatAmI(1)
  whatAmI("hey")
}
```

## Arrays

在Go中，数组是特定长度元素的编号序列。

```go
package main
import "fmt"
func main(){
  //这里创建了一个5个int元素的数组。默认是零值，对于int而言就是0
  var a [5]int
  fmt.Println("emp:",a)
  //可以通过array[index]=value语法设值，或者通过array[index]取值
  a[4]=100
  fmt.Println("set:",a)
  fmt.Println("get:",a[4])
  //内置的len函数返回数组的长度
  fmt.Println("len:",len(a))
  //声明并初始化数组
  b:=[5]int{1,2,3,4,5}
  fmt.Prinln("dc1:",b)
  //数组是一维的，但你可以组合类型来构建多维数组结构
  var twoD [2][3]int
  for i:=0;i<2;i++{
    for j:=0;j<3;j++{
      twoD[i][j]=i+j
    }
  }
  fmt.Println("2d: ",twoD)
}
```

当使用fmt.Println方法打印时，数组将以[v1 v2 v3 ... ]的形式展现

在典型的Go程序中，你将发现slice用的要比array更多，我们将在下一节看到它。

## Slices

slice是一个重要的数据类型，对于序列提供了比数组更强大的接口

```go
package main
import "fmt"
func main(){
  //与数组不同，切片仅由其包含的元素（不是元素的数量）键入。 
  //要创建一个非零长度的空切片，请使用内置的make。 这里我们制作长度为3的字符串（最初为零值）。
  s:=make([]string,3)
  fmt.Println("emp:",s)
  //可以像数组一样set和get
  s[0]="a"
  s[1]="b"
  s[2]="c"
  fmt.Println("set:",s)
  fmt.Println("get:",s[2])
  //len能够返回slice的长度
  fmt.Println("len:",len(s))
  //作为这些基本操作的补充，slice支持一些令其比数组更丰富的东西。
  //其中一个是append函数，其返回一个包含一个或多个新值的slice
  //注意需要接受append的返回值来获取新的slice值
  s=append(s,"d")
  s=append(s,"e","f")
  fmt.Println("apd:",s)
  //slice可以复制
  c:=make([]string,len(s))
  copy(c,s)
  fmt.Println("cpy:",c)
  //slice支持切片操作，语法为slice[low:high]。如下将得到元素s[2],s[3]和s[4]
  l:=s[2:5]
  fmt.Println("sl1:",l)
  //如下切片截止到（不包含）s[5]
  l=s[:5]
  fmt.Pritnln("sl2:",l)
  //如下切片从s[2]开始（包含）
  l=s[2:]
  fmt.Println("sl3:",l)
  //声明并初始化slice
  t:=[]string{"g","h","i"}
  fmt.Println("dcl:",t)
  //slice也可以组织成一个多为数据结构，内部的slice长度可以变化，这与多维数组不同
  twoD:=make([][]int,3)
  for i:=0;i<3;i++{
    innerLen:=i+1
    twoD[i]=make([]int,innerLen)
    for j:=0;j<innerLen;j++{
      twoD[i][j]=i+j
    }
  }
  fmt.Println("2d: ",twoD)
}
```

虽然slice与array是不同的类型，但是使用fmt.Println的展示结果很相似

可以通过这篇Go团队的[文章](https://blog.golang.org/go-slices-usage-and-internals)来获得设计和实现slice的更多细节

现在我们看完了array和slice，下面是另一个重要的内置数据结构：map

## Maps

Map是Go内置的关联数据类型（其他语言可能成为哈希或者字典）

```go
package main
import "fmt"
func main(){
  //要创建一个空的map，使用内置make函数：make(map[key-type]val-type)
  m:=make(map[string]int)
  //通过经典的name[key]=val语法来设置key/value对
  m["k1"]=7
  m["k2"]=13
  fmt.Println("map:",m)
  //通过name[key]来获取一个value
  v1:=m["k1"]
  fmt.Println("v1: ",v1)
  //len函数返回map中键值对的个数
  fmt.Println("len:",len(m))
  //内置的delete函数将移除map中的键值对
  delete(m,"k2")
  fmt.Println("map:",m)
  //第一个值是该key的value，但此处不需要，故使用空白占位符“_”忽略
  //第二个可选的返回值表明该键是否在map中，这样可以消除不存在的键，和键值为0或者""的歧义
  _,prs:=m["k2"]
  fmt.Println("prs:",prs)
  //声明并初始化map
  n:=map[string]int{"foo":1,"bar":2}
  fmt.Println("map:",n)
}
```

map将以[ k:v k:v ]的形式打印

## Range

range可以遍历各种数据结构中的元素。让我们看看range是如何在我们已经学过的数据结构中应用的。

```go
package main
import "fmt"
func main(){
  //这里使用range来计算元素的和，数组也是类似的用法
  nums:=[]int{2,3,4}
  sum:=0
  for _,num:=range nums{
    sum+=num
  }
  fmt.Println("sum:",sum)
  //在slice和array上的range均为每个条目提供了索引和值
  for i,num:=range nums{
    if num==3{
      fmt.Pritnln("index",i)
    }
  }
  //map上的range通过key/value对进行遍历
  kvs:=map[string]string{"a":"apple","b":"banana"}
  for k,v:=range kvs{
    fmt.Printf("%s -> %s\n",k,v)
  }
  //range可以仅通过key进行遍历
  for k:= range kvs{
    fmt.Println("key:",k)
  }
  //range作用在string上将得到unicode code points，第一个值是字符的起始字节索引，第二个值是字符本身
  for i,c:range "go" {
    cmt.Println(i,c)
  }
}
```

## Functions

函数是Go的核心内容。我们将通过不同的例子来学习函数的相关内容。

```go
package main
import "fmt"
//这个函数接收两个int并以int类型返回他们的和
func plus(a int,b int) int{
  //Go需要显式的return语句，它不会自动返回最后一个表达式的值
  return a+b
}
//如果有多个连续的相同类型的参数，可以忽略前面的类型声明
func plusPlus(a,b,c int) int{
  return a+b+c
}
func main(){
  //如你所想的那样调用函数
  res := plus(1,2)
  fmt.Println("1+2=",res)
  res = plusPlus(1,2,3)
  fmt.Println("1+2+3=",res)
}
```

Go函数还有些其他的特色，其中一个是返回多个值，我们在下一节看。

## Multiple Return Values

Go内置支持了返回多个值。这一特点经常用于Go的习惯用法，例如同时返回结果和错误值

```go
package main
import "fmt"
//方法签名中的(int,int)表明它将返回两个int值
func vals()(int,int){
  return 3,7
}
func main(){
  //这里我们通过多重赋值来使用两个不同的返回值
  a,b := vals()
  fmt.Println(a)
  fmt.Println(b)
  //如果只需要返回结果的子集，使用空白占位符_
  _,c := vals()
  fmt.Println(c)
}
```

接受不同个数的参数是Go函数的另一个很棒的特点，我们将在下一节看到。

## Variadic Functions

变参函数，可以使用任意数量的参数来进行调用。例如fmt.Println是一个常见的变参函数。

```go
package main
import "fmt"
//这是一个接收任意数量的int值的函数
func sum(nums ...int){
  fmt.Print(nums," ")
  total := 0
  for _, num := range nums{
    total += num
  }
  fmt.Println(total)
}
func main(){
  sum(1,2)
  sum(1,2,3)
  //如果你已经在一个slice中定义了多个参数，可以使用func(slice...)来直接应用到变参函数中
  nums :=[]int{1,2,3,4}
  sum(nums...)
}
```

关于Go函数另一个关键方面是它形成闭包的能力，下一节可以看到。

## Closures

Go支持匿名函数，它可以形成闭包。当你想要定义一个不记名的内部函数时，匿名函数就很有用了。

```go
package main
import "fmt"
//intSeq函数返回另一个函数，它定义在了intSeq函数内部，并且是匿名的。
//返回的函数关闭变量i以形成闭包
func intSeq() func int{
  i := 0
  return func() int{
    i+=1
    return i
  }
}
func main(){
  //调用intSeq，并将结果(一个函数)赋予nextInt
  //这个函数持有一个自己的i值，每次调用nextInt时都会更新
  nextInt:=intSeq()
  //多次调用nextInt函数可以看到闭包的效果
  //zephyr:如非闭包写法，每次函数都会进行初始化变量i，反复调用intSeq不会有这个效果
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Pritnln(nextInt())
  //要确认该状态对该特定函数是唯一的，请创建并测试一个新函数。
  newInts:=intSeq()
  fmt.Println(newInts())
}
```

函数的最后一个特点是递归。

## Recursion

Go支持递归函数。下面是一个经典的斐波那契数列示例

```go
package main
import "fmt"
func fact(n int) int {
  if n==0 {
    return 1
  }
  //fact函数调用自身，直到n为0
  return n * fact(n-1)
}
func main(){
  fmt.Println(fact(7))
}
```

下一节：指针。

## Pointers

```go
package main
import "fmt"
func zeroval(ival int){
  ival = 0
}
func zeroptr(iptr *int){
  *iptr = 0
}
func main(){
  i:=1
  fmt.Println("initial:",i)
  zeroval(i)
  fmt.Println("zeroval:",i)
  zeroptr(&i)
  fmt.Println("zeroptr:",i)
  fmt.Println("pointer:",&i)
}
```

## Structs

```go
package main
import "fmt"
type person struct{
  name string
  age int
}
func main(){
  fmt.Println(person{"Bob",20})
  fmt.Println(person{name:"Alice",age:30})
  fmt.Println(person{name:"Fred"})
  fmt.Println(&person{name:"Ann",age:40})
  s:=person{name:"Sean",age:50}
  sp:=&s
  fmt.Println(sp.age)
  sp.age=51
  fmt.Println(sp.age)
}
```

## Methods

```go
package main
import "fmt"
type rect struct{
  width,height int
}
func (r *rect) area() int{
  return r.width * r.height
}
func (r rect) perim() int{
  return 2*r.width+2*r.height
}
func main(){
  r:=rect{width:10,height:5}
  fmt.Println("area: ",r.area())
  fmt.Println("perim: ",r.perim())
  rp:=&r
  fmt.Println("area: ",rp.area())
  fmt.Println("perim: ",rp.perim())
}
```

## Interfaces

```go
package main
import "fmt"
import "math"
type geometry interface{
  area() float64
  perim() float64
}
type rect struct{
  width,height float64
}
type circle struct{
  radius float64
}
func (r rect) area() float64{
  return r.width * r.height
}
func (r rect) perim() float64{
  return 2*r.width+2*r.height
}
func (c circle) area() float64{
  return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64{
  return 2 * math.Pi * c.radius
}
func measure(g geometry){
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.perim())
}
func main(){
  r:=rect{width:3,height:4}
  c:=circle{radius:5}
  measure(r)
  measure(c)
}
```

## Errors

```go
package main
import "errors"
import "fmt"
fuc f1(arg int)(int,error){
  if arg==42{
    return -1,errors.New("can't work with 42")
  }
  return arg+3,nil
}
type argError struct{
  arg int
  prob string
}
type argError struct{
  arg int
  prob string
}
func (e *argError) Error() string{
  return fmt.Sprintf("%d - %s",e.arg,e.prob)
}
func f2(arg int)(int,error){
  if arg==42{
    return -1,&argError{arg,"can't work with it"}
  }
  return arg+3,nil
}
func main(){
  for _,i:=range []int{7,42}{
    if r,e:=f1(i);e!=nil{
      fmt.Println("f1 failed:",e)
    }else{
      fmt.Println("f1 worked:",r)
    }
  }
  _,e:=f2(42)
  if ae,ok := e.(*argError);ok{
    fmt.Println(ae.arg)
    fmt.Println(ae.prob)
  }
}
```

## Goroutines

```go
package main
import "fmt"
func f(from string){
  for i:=0;i<3;i++{
    fmt.Println(from,";",i)
  }
}
func main(){
  f("direct")
  go f("goroutine")
  go func(msg string){
	fmt.Println(msg)
  }("going")
  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}
```

## Channels

```go
package main
import "fmt"
func main(){
  messages:=make(chan string)
  go func(){messages<-"ping"}()
  msg:=<-messages
  fmt.Println(msg)
}
```

## Channel Buffering

```go
package main
import "fmt"
func main(){
  messages:=make(chan string,2)
  messages<-"buffered"
  messages<-"channel"
  fmt.Println(<-messages)
  fmt.Println(<-messages)
}
```

## Channel Synchronization

```go
package main
import "fmt"
import "time"
func worker(donw chan bool){
  fmt.Print("working...")
  time.Sleep(time.Second)
  fmt.Println("done")
  done<-true
}
func main(){
  done:=make(chan bool,1)
  go worker(done)
  <-done
}
```

## Channel Directions

```go
package main
import "fmt"
func ping(pings chan<- string,msg string){
  pings<-msg
}
func pong(pings <-chan string,pongs chan<- string){
  msg:=<-pings
  pongs<-msg
}
func main(){
  pings:=make(chan string,1)
  pongs:=make(chan string,1)
  ping(pings,"passed message")
  pong(pings,pongs)
  fmt.Println(<-pongs)
}
```

## Select

```go
package main
import "time"
import "fmt"
func main(){
  c1:=make(chan string)
  c2:=make(chan string)
  go func(){
    time.Sleep(time.Second*1)
    c1<-"one"
  }()
  go func(){
    time.Sleep(time.Second*2)
    c2<-"two"
  }()
  for i:=0;i<2;i++{
    select{
    case msg1:=<-c1:
      fmt.Println("received",msg1)
    case msg2:=<-c2:
      fmt.Println("received",msg2)
    }
  }
}
```

## Timeouts

```go
package main
import "time"
import "fmt"
func main(){
  c1:=make(chan string,1)
  go func(){
    time.Sleep(time.Second*2)
    c1<-"result 1"
  }()
  select{
  case res:=<-c1:
    fmt.Println(res)
  case <-time.After(time.Second*1):
  	fmt.Println("timeout 1")  
  }
  c2:=make(chan string,1)
  go func(){
    time.Sleep(time.Second*2)
    c2<-"result 2"
  }()
  select{
  case res:=<-c2:
    fmt.Println(res)
  case <-time.After(time.Second*3):
    fmt.Printnln("timeout 2")
  }
}
```

## Non-Blocking Channel Operations

```go
package main
import "fmt"
func main(){
  messages:=make(chan string)
  signals:=make(chan bool)
  select{
  case msg:=<-messages:
    fmt.Println("received message",msg)
  default:
    fmt.Pritln("no message received")
  }
  msg:="hi"
  select{
  case messages<-msg:
    cmt.Println("sent message",msg)
  default:
    fmt.Println("no message sent")
  }
  select{
  case msg:=<-messages:
    fmt.Println("received message",msg)
  case sig:=<-signals:
    fmt.Println("received signal",sig)
  default:
    fmt.Println("no activity")
  }
}
```

## Closing Channels

```go
package main
import "fmt"
func main(){
  jobs:=make(chan int,5)
  done:=make(chan bool)
  go func(){
    for{
      j,more:=<-jobs
      if more{
        fmt.Println("received job",j)
      }else{
        fmt.Println("received all jobs")
        done<-true
        return
      }
    }
  }()
  for j:=1;j<=3;j++{
    jobs<-j
    fmt.Println("sent job",j)
  }
  close(jobs)
  fmt.Println("sent all jobs")
  <-done
}
```

## Range over Channels

```go
package main
import "fmt"
func main(){
  queue:=make(chan string,2)
  queue<-"one"
  queue<-"two"
  close(queue)
  for elem:=range queue{
    fmt.Println(elem)
  }
}
```

## Timers

```go
package main
import "time"
import "fmt"
func main(){
  timer1:=time.NewTimer(time.Second*2)
  <-timer1.C
  fmt.Println("Timer 1 expired")
  timer2:=time.NewTimer(time.Second)
  go func(){
    <-timer2.C
    fmt.Println("Timer 2 expired")
  }()
  stop2:=timer2.Stop()
  if stop2{
    fmt.Println("Timer 2 stopped")
  }
}
```

## Tickers

```go
package main
import "time"
import "fmt"
func main(){
  ticker:=time.NewTicker(time.Millisecond*500)
  go func(){
    for t:=range ticker.C{
	  fmt.Prinltn("Tick at ",t)
    }
  }()
  time.Sleep(time.Millisecond*1600)
  ticker.Stop()
  fmt.Println("Ticker stopped")
}
```

## Worker Pools

```go
package main
import "fmt"
import "time"
func worker(id int,jobs <-chan int,results chan<- int){
  for j:=range jobs{
    fmt.Println("worker",id,"started job",j)
    time.Sleep(time.Second)
    fmt.Println("worker",id,"finished job",j)
    results <- j*2
  }
}
func main(){
  jobs:=make(chan int,100)
  results:=make(chan int,100)
  for w:=1;w<=3;w++{
    go worker(w,jobs,results)
  }
  for j:=1;j<=5;j++{
    jobs<-j
  }
  close(jobs)
  for a:=1;a<=5;a++{
    <-results
  }
}
```

## Rate Limiting

```go
package main
import "time"
import "fmt"
func main(){
  requests:=make(chan int,5)
  for i:=1;i<=5;i++{
    request<-i
  }
  close(requests)
  limiter:=time.Tick(time.Millisecond*200)
  for req:=range request{
    <-limiter
    fmt.Println("request",req,time.now())
  }
  burstyLimiter:=make(chan time.Time,3)
  for i:=0;i<3;i++{
    burstyLimiter<-time.Now()
  }
  go func(){
    for t:=range time.Tick(time.Millisecond*200){
      burstyLimiter<-t
    }
  }()
  burstyRequests:=make(chan int,5)
  for i:=1;i<=5;i++{
    burstyRequest<-i
  }
  close(burstyRequests)
  for req:=range burstyRequests{
    <-burstyLimiter
    fmt.Println("request",req,time.Now())
  }
}
```

## Atomic Counters

```go
package main
import "fmt"
import "time"
import "sync/atomic"

func main(){
  var ops unit64=0
  for i:=0;i<50;i++{
    go func(){
      for{
        atomic.AddUnit64(&ops,1)
        time.Sleep(time.Millisecond)
      }
    }()
  }
  time.Sleep(time.Second)
  opsFinal:=atomic.LoadUnit64(&ops)
  fmt.Println("ops:",opsFinal)
}
```

## Mutexes

```go
package main
import(
  "fmt"
  "math/rand"
  "sync"
  "sync/atomic"
  "time"
)
func main(){
  var state=make(map[int]int)
  var mutex=&sync.Mutex{}
  var readOps unit64=0
  var writeOps unit64=0
  for r:=0;r<100;r++{
    go func(){
      total:=0
      for{
        key:=rand.Intn(5)
        mutex.Lock()
        total+=state[key]
        mutex.Unlock()
        atomic.AddUnit64(&readOps,1)
        time.Sleep(time.Millisecond)
      }
    }()
  }
  for w:=0;w<10;w++{
    go func(){
      for{
        key:=rand.Intn(5)
        val:=rand.Intn(100)
        mutex.Lock()
        state[key]=val
        mutex.Unlock()
        atomic.AddUnit64(&writeOps,1)
        time.Sleep(time.Millisecond)
      }
    }()
  }
  time.Sleep(time.Second)
  readOpsFinal:=atomic.LoadUnit64(&readOps)
  fmt.Println("readOps:",readOpsFinal)
  writeOpsFinal:=atomic.LoadUnit64(&writeOps)
  fmt.Println("writeOps:",writeOpsFinal)
  mutex.Lock()
  fmt.Println("state:",state)
  mutex.Unlock()
}
```

## Stateful Goroutines

```go
package main
import(
  "fmt"
  "math/rand"
  "sync/atomic"
  "time"
)
type readOp struct{
  key int
  resp chan int
}
type writeOp struct{
  key int
  val int
  resp chan bool
}
func main() {
  var readOp unit64=0
  var writeOps unit64=0
  reads:=make(chan *readOp)
  writes:=make(chan *writeOp)
  go func(){
    var state=make(map[int]int)
    for{
      select{
      case read:=<-reads:
        read.resp<-state[read.key]
      case write:=<-writes:
        state[write.key]=write.val
        write.resp<-true
      }
    }
  }()
  for r:=0;r<100;r++{
    go func(){
      for {
        read:=&readOp{
          key:rand.Intn(5),
          resp:make(chan int)
        }
        reads<-read
        <-read.resp
        atomic.AddUnit64(&readOps,1)
        time.Sleep(time.Millisecond)
      }
    }()
  }
  for w:=0;w<10;w++{
    go func(){
      for {
        write:=&writeOp{
          key:rand.Intn(5),
          val:rand.Intn(100),
          resp:make(chan bool)
        }
      }
    }()
  }
  time.Sleep(time.Second)
  readOpsFinal:=atomic.LoadUnit64(&readOps)
  fmt.Println("readOps:",readOpsFinal)
  writeOpsFinal:=atomic.LoadUnit64(&writeOps)
  fmt.Println("writeOps:",writeOpsFinal)
}
```

## Sorting

```go
package main
import "fmt"
import "sort"
func main(){
  strs:=[]string{"c","a","b"}
  sort.Strings(strs)
  fmt.Println("Strings:",strs)
  ints:=[]int{7,2,4}
  sort.Ints(ints)
  fmt.Println("Ints:",ints)
  s:=sort.IntsAreSorted(ints)
  fmt.Println("Sorted:",s)
}
```

## Sorting by Functions

```go
package main
import "fmt"
import "sort"
type ByLength []string
func (s ByLength) Len() int{
  return len(s)
}
func (s ByLength) Swap(i,j int){
  s[i],s[j] = s[j],s[i]
}
func (s ByLength) Less(i,j int) bool {
  return len(s[i])<len(s[j])
}
func main(){
  fruits:=[]string{"peach","banana","kiwi"}
  sort.Sort(ByLength(fruits))
  fmt.Println(fruits)
}
```

## Panic

```go
package main
import "os"
func main(){
  panic("a problem")
  _,err:=os.Create("/tmp/file")
  if err!=nil{
    panic(err)
  }
}
```

## Defer

```go
packgae main
import "fmt"
import "os"
func main(){
  f:=createFile("/tmp/defer.txt")
  defer closeFile(f)
  writeFile(f)
}
func createFile(p string) *os.File{
  fmt.Println("creating")
  f,err:=os.Create(p)
  if err!=nil{
    panic(err)
  }
  return f
}
func writeFile(f *os.File){
  fmt.Println("writing")
  fmt.Println(f,"data")
}
func closeFile(f *os.File){
  fmt.Println("closing")
  f.Close()
}
```
## Collection Functions
```go
package main
import "strings"
import "fmt"
func Index(vs []string,t string) int{
  for i,v:=range vs{
    if v==t{
      return i
    }
  }
  return -1
}
func Include(vs []string,t string) bool{
  return Index(vs,t)>=0
}
func Any(vs []string,f func(string) bool) bool{
  for _,v:=range vs{
    if f(v){
      return true
    }
  }
  return false
}
func All(vs []string,f func(string) bool) bool{
  for _,v:=range vs{
    if !f(v){
      return false
    }
  }
  return true
}
func Filter(vs []string,f func(string) bool) []string{
  vsf :=make([]string,0)
  for _,v:=range vs{
    if f(v){
      vsf=append(vsf,v)
    }
  }
  return vsf
}
func main(){
  var strs=[]string{"peach","apple","pear","plum"}
  fmt.Println(Index(strs,"pear"))
  fmt.Println(Include(strs,"grape"))
  fmt.Println(Any(strs,func(v string) bool{
    return strings.HasPrefix(v,"p")
  }))
  fmt.Println(All(strs,func(v string) bool{
    return strings.HasPrefix(v,"p")
  }))
  fmt.Println(Filter(strs,func(v string) bool{
    return strings.Contains(v,"e")
  }))
  fmt.Println(Map(strs,strings.ToUpper))
}
```

## String Functions

```go
package main
import s "strings"
import "fmt"
var p=fmt.Println
func main(){
  p("Contains: ",s.Contains("test","es"))
  p("Count: ",s.Count("test","t"))
  p("HasPrefix: ",s.HasPrefix("test","te"))
  p("HasSuffix: ",s.HasSuffix("test","st"))
  p("Index: ",s.Index("test","e"))
  p("Join: ",s.Join([]string{"a","b"},"-"))
  p("Repeat: ",s.Repeat("a",5))
  p("Replace: ",s.Replace("foo","o","0",-1))
  p("Replace ",s.Replace("foo","o","0",1))
  p("Split: ",s.Split("a-b-c-d-e","-"))
  p("ToLower: ",s.ToLower("TEST"))
  p("ToUpper: ",s.ToUpper("test"))
  p()
  p("Len: ",len("hello"))
  p("Char: ","hello"[1])
}
```

## String Formatting

```go
package main
import "fmt"
import "os"
type point struct{
  x,y int
}
func main(){
  p:=point{1,2}
  fmt.Printf("%v\n",p)
  fmt.Printf("%+v\n",p)
  fmt.Printf("%#v\n",p)
  fmt.Printf("%T\n",p)
  fmt.Printf("%t\n",true)
  fmt.Printf("%d\n",123)
  fmt.Printf("%b\n",14)
  fmt.Printf("%c\n",33)
  fmt.Printf("%x\n",456)
  fmt.Printf("%f\n",78.9)
  fmt.Printf("%e\n",123400000.0)
  fmt.Printf("%E\n",123400000.0)
  fmt.Pirntf("%s\n","\"string\"")
  fmt.Printf("%q\n","\"string\"")
  fmt.Printf("%x\n","hex this")
  fmt.Printf("%p\n",&p)
  fmt.Printf("|%6d|%6d|\n",12,345)
  fmt.Printf("|%6.2f|%6.2f|\n",1.2,3.45)
  fmt.Printf("|%-6.2f|%-6.2f|\n",1.2,3.45)
  fmt.Printf("|%6s|%6s|\n","foo","b")
  fmt.Printf("|%-6s|%-6s|\n","foo","b")
  s:=fmt.Sprintf("a %s","string")
  fmt.Println(s)
  fmt.Fprintf(os.Stderr,"an %s\n","error")
}
```

## Regular Expressions

```go
package main
import "bytes"
import "fmt"
import "regexp"
func main(){
  math,_:=regexp.MatchString("p([a-z]+)ch","peach")
  fmt.Printf(match)
  r,_:=regexp.Compile("p([a-z]+)ch")
  fmt.Println(r.MatchString("peach"))
  fmt.Println(r.FindString("peach punch"))
  fmt.Println(r.FindStringIndex("peach punch"))
  fmt.Println(r.FindStringSubmatch("peach punch"))
  fmt.Println(r.FindStringSubmatchIndex("peach punch"))
  fmt.Println(r.FindAllString("peach punch pinch",-1))
  fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch",-1))
  fmt.Println(r.FindAllString("peach punch pinch",2))
  fmt.Println(r.Match([]byte("peach")))
  r=regexp.MustComplie("p([a-z]+)ch")
  fmt.Println(r)
  fmt.Println(r.ReplaceAllString("a peach","<fruit>"))
  in:=[]byte("a peach")
  out:=r.ReplaceAllFunc(in,bytes.ToUpper)
  fmt.Println(string(out))
}
```

## JSON

```go
package main
import "encoding/json"
import "fmt"
import "os"
type Response1 struct{
  Page int
  Fruits []string
}
type Response2 struct{
  Page int `json:"page"`
  Fruits []string `json:"fruits"`
}
func main(){
  bolB,_:=json.Marshal(true)
  fmt.Println(string(bolB))
  intB,_:=json.Marshal(1)
  fmt.Println(string(intB))
  fltB,_:=json.Marshal(2.34)
  fmt.Println(string(fltB))
  strB,_:=json.Marshal("gopher")
  fmt.Println(string(strB))
  slcD:=[]string{"apple","peach","pear"}
  slcB,_:=json.Marshal(slcD)
  fmt.Println(string(slcB))
  mapD:=map[string]int{"apple":5,"lettuce":7}
  mapB,_:=json.Marshal(mapD)
  fmt.Println(string(mapB))
  res1D:=&Response1{
    Page:1,
    Fruits:[]string{"apple","peach","pear"}
  }
  res1B,_:=json.Marshal(res1D)
  fmt.Println(string(res1B))
  res2D:=&Reponse2{
    Page:1,
    Fruits:[]string{"apple","peach","pear"}
  }
  res2B,_:=json.Marshal(res2D)
  fmt.Println(string(res2B))
  byt:=[]byte(`{"num":6.13,"strs":["a","b"]}`)
  var dat map[string]interface{}
  if err:=json.Unmarshal(byt,&dat);err!=nil{
    panic(err)
  }
  fmt.Println(dat)
  num:dat["num"].(float64)
  fmt.Println(num)
  strs:=dat["strs"].([]interface{})
  str1:=strs[0].(string)
  fmt.Println(str1)
  str:=`{"page":1,"fruits":["apple","peach"]}`
  res:=Response2{}
  json.Unmarshal([]byte(str),&res)
  fmt.Println(res)
  fmt.Println(res.Fruits[0])
  enc:json.NewEncoder(os.Stdout)
  d:=map[string]int{"apple":5,"lettuce":7}
  enc.Encode(d)
}
```

## Time

```go
package main
import "fmt"
import "time"
func main(){
  p:=fmt.Pritnln
  now:=time.Now()
  p(now)
  then:=time.Date(2009,11,17,20,34,58,651387237,time.UTC)
  p(then)
  p(then.Year())
  p(then.Month())
  p(then.Day())
  p(then.Hour())
  p(then.Minute())
  p(then.Second())
  p(then.Nanosecond())
  p(then.Location())
  p(then.Weekday())
  p(then.Before(now))
  p(then.After(now))
  p(then.Equal(now))
  diff:=now.Sub(then)
  p(diff)
  p(diff.Hours())
  p(diff.Minutes())
  p(diff.Seconds())
  p(diff.Nanoseconds())
  p(then.Add(diff))
  p(then.Add(-diff))
}
```

## Epoch

```go
package main
import "fmt"
import "time"
func main(){
  now:=time.Now()
  secs:=now.Unix()
  nanos:=now.UnixNano()
  fmt.Println(now)
  millis:=nanos/1000000
  fmt.Println(secs)
  fmt.Println(millis)
  fmt.Println(nanos)
  fmt.Println(time.Unix(secs,0))
  fmt.Println(time.Unix(0,nanos))
}
```

## Time Formatting / Parsing

```go
package main
import "fmt"
import "time"
func main(){
  p:=fmt.Println
  t:=time.Now()
  p(t.Format(time.RFC3339))
  t1,e:=time.Parse(
    time.RFC3339,
    "2012-11-01T22:08:41+00:00"
  )
  p(t1)
  p(t.Format("3:04PM"))
  p(t.Format("Mon Jan _2 15:04:05 2006"))
  p(t.Format("2006-01-02T15:04:05.999999-07:00"))
  form:="3 04 PM"
  t2,e:=time.Parse(form,"8 41 PM")
  p(t2)
  fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
    t.Year(),t.Month(),t.Day(),
    t.Hour(),t.Minute(),t.Second()
  )
  ansic:="Mon Jan _2 15:04:05 2006"
  _,e=time.Parse(ansic,"8:41PM")
  p(e)
}

## Random Numbers

​```go
package main
import "time"
import "fmt"
import "math/rand"
func main(){
  fmt.Print(rand.Intn(100),",")
  fmt.Print(rand.Intn(100))
  fmt.Println()
  fmt.Println(rand.Float64())
  fmt.Print((rand.Float64()*5)+5,",")
  fmt.Print((rand.Float64()*5)+5)
  fmt.Println()
  s1:=rand.NewSource(time.Now().UnixNano())
  r1:=rand.New(s1)
  fmt.Print(r1.Intn(100),",")
  fmt.Print(r1.Intn(100))
  fmt.Println()
  s2:=rand.NewSrouce(42)
  r2:=rand.New(s2)
  fmt.Print(r2.Intn(100),",")
  fmt.Print(r2.Intn(100))
  fmt.Println()
  s3:=rand.NewSource(42)
  r3:=rand.New(s3)
  fmt.Print(r3.Intn(100),",")
  fmt.Print(r3.Intn(100))
}

## Number Parsing

​```go
package main
import "strconv"
import "fmt"
func main(){
  f,_:=strconv.ParseFloat("1.234",64)
  fmt.Println(f)
  i,_:=strconv.ParseInt("123",0,64)
  fmt.Println(i)
  d,_:=strconv,Parseint("0x1c8",0,64)
  fmt.Println(d)
  u,_:=strconv.ParseUint("789",0,64)
  fmt.Println(u)
  k,_:=strconv.Atoi("135")
  fmt.Println(k)
  _,e:=strconv.Atoi("wat")
  fmt.Println(e)
}
```

## URL Parsing

```go
package main
import "fmt"
import "net"
import "net/url"
func main(){
  s:="postgres://user:pass@host.com:5432/path?k=v#f"
  u,err:=url.Parse(s)
  if err!=nil{
    panic(err)
  }
  fmt.Println(u.Scheme)
  fmt.Println(u.User)
  fmt.Println(u.User.Username())
  p,_:=u.User.Password()
  fmt.Println(p)
  fmt.Println(u.Host)
  host,port,_:=net.SplitHostPort(u.Host)
  fmt.Println(host)
  fmt.Println(port)
  fmt.Println(u.Path)
  fmt.Println(u.Fragment)
  fmt.Println(u.RawQuery)
  m,_:=url.ParseQuery(u.RawQuery)
  fmt.Println(m)
  fmt.Println(m["k"][0])
}
```

## SHA1 Hashes

```go
package main
import "crypto/sha1"
import "fmt"
func main(){
  s:="sha1 this string"
  h:=sha1.New()
  h.Write([]byte(s))
  bs:=h.Sum(nil)
  fmt.Println(s)
  fmt.Printf("%x\n",bs)
}
```

## Base64 Encoding

```go
package main
import b64 "encoding/base64"
import "fmt"
func main(){
  data:="abc123!?$*&()'-=@~"
  sEnc:=b64.StdEncoding.EncodeToString([]byte(data))
  fmt.Println(sEnc)
  sDec,_:=b64.StdEncoding.DecodeString(sEnc)
  fmt.Println(string(sDec))
  fmt.Println()
  uEnc:=b64.URLEncoding.EncodeToString([]byte(data))
  fmt.Println(uEnc)
  uDec,_:=b64.URLEncoding.DecodeString(uEnc)
  fmt.Println(string(uDec))
}
```

## Reading Files

```go
package main
import(
  "bufio"
  "fmt"
  "io"
  "io/ioutil"
  "os"
)
func check(e error){
  if e!=nil{
    panic(e)
  }
}
func main(){
  dat,err:=ioutil.ReadFile("/tmp/dat")
  check(err)
  fmt.Print(string(dat))
  f,err:=os.Open("/tmp/dat")
  check(err)
  b1:=make([]byte,5)
  n1,err:=f.Read(b1)
  check(err)
  fmt.Printf("%d bytes: %s\n",n1,string(b1))
  o2,err:=f.Seek(6,0)
  check(err)
  b2:=make([]byte,2)
  n2,err:=f.Read(b2)
  check(err)
  fmt.Printf("%d bytes @ %d: %s\n",n2,o2,string(b2))
  o3,err:=f.Seek(6,0)
  check(err)
  b3:=make([]byte,2)
  n3,err:=io.ReadAtLeast(f,b3,2)
  check(err)
  fmt.Printf("%d bytes @ %d: %s\n",n3,o3,string(b3))
  _,err=f.Seek(0,0)
  check(err)
  r4:=bufio.NewReader(f)
  b4,err:=r4.Peek(5)
  check(err)
  fmt.Printf("5 bytes: %s\n",string(b4))
  f.Close()
}
```

## Writing Files

```go
package main
import(
  "bufio"
  "fmt"
  "io/ioutil"
  "os"
)
func check(e error){
  if e!=nil{
    panic(e)
  }
}
func main(){
  d1:=[]byte("hello\ngo\n")
  err:=ioutil.WriteFile("/tmp/dat1",d1,0644)
  check(err)
  f,err:=os.Create("/tmp/dat2")
  check(err)
  defer f.Close()
  d2:=[]byte{115,111,109,101,10}
  n2,err:=f.Write(d2)
  check(err)
  fmt.Printf("wrote %d bytes\n",n2)
  n3,err:=f.WriteString("writes\n")
  fmt.Printf("wrote %d bytes\n",n3)
  f.Sync()
  w:=bufio.NewWriter(f)
  n4,err:=w.WriteString("buffered\n")
  fmt.Printf("wrote %d bytes\n",n4)
  w.Flush()
}
```

## Line Filters

```go
package main
import (
  "bufio"
  "fmt"
  "os"
  "strings"
)
func main(){
  scanner:=bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    ucl:=strings.ToUpper(scanner.Text())
    fmt.Println(ucl)
  }
  if err:=scanner.Err();err!=nil{
    fmt.Fprintln(os.Stderr,"error:",err)
    os.Exit(1)
  }
}
```

## Command-Line Arguments

```go
package main
import "os"
import "fmt"
func main(){
  argsWithProg:=os.Args
  argsWithoutProg:=os.Args[1:]
  arg:=os.Args[3]
  fmt.Println(argsWithProg)
  fmt.Println(argsWithoutProg)
  fmt.Println(arg)
}
```
## Command-Line Flags

```go
package main
import "flag"
import "fmt"
func main(){
  wordPtr:=flag.String("word","foo","a string")
  numbPtr:=flag.Int("numb",42,"an int")
  boolPtr:=flag.Bool("fork",false,"a bool")
  var svar string
  flag.StringVar(&svar,"svar","bar","a string var")
  flag.Parse()
  fmt.Println("word:",*wordPtr)
  fmt.Println("numb:",*numbPtr)
  fmt.Println("fork:",*boolPtr)
  fmt.Println("svar:",svar)
  fmt.Println("tail:",flag.Args())
}
```

## Environment Variables

```go
package main
import "os"
import "strings"
import "fmt"
func main(){
  os.Setenv("FOO","1")
  fmt.Println("FOO:",os.Getenv("FOO"))
  fmt.Println("BAR:",os.Getenv("BAR"))
  fmt.Println()
  for _,e:=range os.Environ(){
    pair:=strings.Split(e,"=")
    fmt.Prinln(pair[0])
  }
}
```

## Spawning Processes
package main
import "fmt"
import "io/iouitl"
import "os/exec"
func main(){
  dateCmd:=exec.Command("date")
  dateOut,err:=dateCmd.Output()
  if err!=nil{
    panic(err)
  }
  fmt.Println("> date")
  fmt.Println(string(dateOut))
  grepCmd:=exec.Command("grep","hello")
  grepIn,_:=grepCmd.StdinPipe()
  grepOut,_:=grepCmd.StdoutPipe()
  grepCmd.Start()
  grepIn.Write([]byte("hello grep\ngoodbye grep"))
  grepIn.Close()
  grepByte,_:=ioutil.ReadAll(grepOut)
  grepCmd.Wait()
  fmt.Println("> grep hello")
  fmt.Println(string(grepBytes))
  lsCmd:=exec.Command("bash","-c","ls -a -l -h")
  lsOut,err:=lsCmd.Output()
  if err!=nil{
    panic(err)
  }
  fmt.Println("> ls -a -l -h")
  fmt.Pritnln(string(lsOut))
}

## Exec'ing Processes

```go
package main
import "syscall"
import "os"
import "os/exec"
func main(){
  binary,lookErr:=exec.LookPath("ls")
  if lookErr!=nil{
    panic(lookErr)
  }
  args:=[]string{"ls","-a","-l","-h"}
  env:=os.Environ()
  execErr:=syscall.Exec(binary,args,env)
  if execErr!=nil{
    panic(execErr)
  }
}
```

## Signals

```go
package main
import "fmt"
import "os"
import "os/signal"
import "syscall"
func main(){
  sigs:=make(chan os.Signal,1)
  done:=make(chan bool,1)
  signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)
  go func(){
    sig:=<-sigs
    fmt.Println()
    fmt.Println(sig)
    done<-true
  }()
  fmt.Println("awaiting signal")
  <-done
  fmt.Println("exiting")
}
```

## Exit

```go
package main
import "fmt"
import "os"
func main(){
  defer fmt.Println("!")
  os.Exit(3)
}
```