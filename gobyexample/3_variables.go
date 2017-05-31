package gobyexample
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
