package gobyexample
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
	fmt.Println(int64(d))
	//数值在使用环境上下文中会得到类型，如变量赋值或者方法调用，如math.Sin需要的是一个float64
	fmt.Println(math.Sin(n))
}