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

```go
package main
import "fmt"
func main(){
  fmt.Println("go"+"lang")
  fmt.Println("1+1=",1+1)
  fmt.Pritnln("7.0/3.0=",7.0/3.0)
  fmt.Println(true&&false)
  fmt.Println(true||false)
  fmt.Pritnln(!true)
}
```

## Variables

```go
package main
import "fmt"
func main(){
  var a string="initial"
  fmt.Println(a)
  var b,c int = 1,2
  fmt.Println(b,c)
  var d = true
  fmt.Println(d)
  var e int
  fmt.Println(e)
  f:="short"
  fmt.Println(f)
}
```

## Constants

```go
package main
import "fmt"
import "math"
const s string = "constant"
func main(){
  fmt.Println(s)
  const n = 50000000
  const d = 3e20 / n
  fmt.Println(d)
  fmt.Prinln(int64(d))
  fmt.Println(math.Sin(n))
}
```

## For

```go
package main
import "fmt"
func main(){
  i := 1
  for i<=3{
    fmt.Println(i)
    i = i+1
  }
  for j:=7;j<=9;j++{
    fmt.Println(j)
  }
  for{
    fmt.Println("loop")
    break
  }
}
```

## If/Else

```go
package main
import "fmt"
func main(){
  if 7%2==0{
    fmt.Println("7 is even")
  }else{
    fmt.Println("7 is odd")
  }
  if 8%4==0{
    fmt.Println("8 is divisible by 4")
  }
  if num:=9;num<0{
    fmt.Println(num,"is negative")
  }else if num<10{
    fmt.Println(num,"has 1 digit")
  }else{
    fmt.Println(num,"has multiple digits")
  }
}
```

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
  switch time.Now().weekday(){
  case time.Saturday,time.Sunday:
    fmt.Println("it's the weekend")
  default:
    fmt.Println("it's a weekday")
  }
  t:=time.Now()
  switch{
  case t.Hour()<12:
    fmt.Println("it's before noon")
  default:
    fmt.Println("it's after noon")
  }
}
```

## Arrays

```go
package main
import "fmt"
func main(){
  var a [5]int
  fmt.Println("emp:",a)
  a[4]=100
  fmt.Println("set:",a)
  fmt.Println("get:",a[4])
  fmt.Println("len:",len(a))
  b:=[5]int{1,2,3,4,5}
  fmt.Prinln("dc1:",b)
  var twoD [2][3]int
  for i:=0;i<2;i++{
    for j:=0;j<3;j++{
      twoD[i][j]=i+j
    }
  }
  fmt.Println("2d: ",twoD)
}
```

## Slices

```go
package main
import "fmt"
func main(){
  s:=make([]string,3)
  fmt.Println("emp:",s)
  s[0]="a"
  s[1]="b"
  s[2]="c"
  fmt.Println("set:",s)
  fmt.Println("get:",s[2])
  fmt.Println("len:",len(s))
  s=append(s,"d")
  s=append(s,"e","f")
  fmt.Println("apd:",s)
  c:=make([]string,len(s))
  copy(c,s)
  fmt.Println("cpy:",c)
  l:=s[2:5]
  fmt.Println("sl1:",l)
  l=s[:5]
  fmt.Pritnln("sl2:",l)
  l=s[2:]
  fmt.Println("sl3:",l)
  t:=[]string{"g","h","i"}
  fmt.Println("dcl:",t)
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

## Maps

```go
package main
import "fmt"
func main(){
  m:=make(map[string]int)
  m["k1"]=7
  m["k2"]=13
  fmt.Println("map:",m)
  v1:=m["k1"]
  fmt.Println("v1: ",v1)
  fmt.Println("len:",len(m))
  delete(m,"k2")
  fmt.Println("map:",m)
  _,prs:=m["k2"]
  fmt.Println("prs:",prs)
  n:=map[string]int{"foo":1,"bar":2}
  fmt.Println("map:",n)
}
```

## Range

```go
package main
import "fmt"
func main(){
  nums:=[]int{2,3,4}
  sum:=0
  for _,num:=range nums{
    sum+=num
  }
  fmt.Println("sum:",sum)
  for i,num:=range nums{
    if num==3{
      fmt.Pritnln("index",i)
    }
  }
  kvs:=map[string]string{"a":"apple","b":"banana"}
  for k,v:=range kvs{
    fmt.Printf("%s -> %s\n",k,v)
  }
  for i,c:range "go" {
    cmt.Println(i,c)
  }
}
```

## Functions

```go
package main
import "fmt"
func plus(a int,b int) int{
  return a+b
}
func plusPlus(a,b,c int) int{
  return a+b+c
}
func main(){
  res := plus(1,2)
  fmt.Println("1+2=",res)
  res = plusPlus(1,2,3)
  fmt.Println("1+2+3=",res)
}
```

## Multiple Return Values

```go
package main
import "fmt"
func vals()(int,int){
  return 3,7
}
func main(){
  a,b := vals()
  fmt.Println(a)
  fmt.Println(b)
  _,c := vals()
  fmt.Println(c)
}
```

## Variadic Functions

```go
package main
import "fmt"
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
  nums :=[]int{1,2,3,4}
  sum(nums...)
}
```

## Closures

```go
package main
import "fmt"
func intSeq() func int{
  i := 0
  return func() int{
    i+=1
    return i
  }
}
func main(){
  nextInt:=intSeq()
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Pritnln(nextInt())
  newInts:=intSeq()
  fmt.Println(newInts())
}
```

## Recursion

```go
package main
import "fmt"
func fact(n int) int {
  if n==0 {
    return 1
  }
  return n * fact(n-1)
}
func main(){
  fmt.Println(fact(7))
}
```

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