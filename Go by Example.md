# Go by Example

## Hello World

```go
package main
import "fmt"
func main(){
  fmt.Println("hello world")
}
```

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