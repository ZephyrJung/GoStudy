package gobyexample
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