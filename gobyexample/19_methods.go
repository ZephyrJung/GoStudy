package gobyexample
import "fmt"
type rect struct{
	width,height int
}
//area方法有一个rect指针类型的接收器
func (r *rect) area() int{
	return r.width * r.height
}
//即可以定义指针类型的接收器，也可以定义值类型的接收器
func (r rect) perim() int{
	return 2*r.width+2*r.height
}
func main(){
	r:=rect{width:10,height:5}
	//调用定义的方法
	fmt.Println("area: ",r.area())
	fmt.Println("perim: ",r.perim())
	//Go为方法调用自动处理了值和引用的转换。使用指针接收器可以避免获得方法调用的拷贝(?)或允许方法修改接收到的struct值
	rp:=&r
	fmt.Println("area: ",rp.area())
	fmt.Println("perim: ",rp.perim())
}
