package gobyexample
import "fmt"
//intSeq函数返回另一个函数，它定义在了intSeq函数内部，并且是匿名的。
//返回的函数关闭变量i以形成闭包
func intSeq() func() int{
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
	fmt.Println(nextInt())
	//要确认该状态对该特定函数是唯一的，请创建并测试一个新函数。
	newInts:=intSeq()
	fmt.Println(newInts())
}
