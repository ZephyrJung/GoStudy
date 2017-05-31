package gobyexample
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
