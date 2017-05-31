package gobyexample
import "fmt"
import "math"
//这是一个geometry的基本接口，本例中将在rect类型和circle类型中实现这个接口
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
//在Go中，实现一个接口只需要实现其中定义的所有方法即可
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
//如果变量是接口类型，那么它可以调用接口内定义的方法
//这是一个通用的measure方法，利用它能够工作在任何geometry上
func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main(){
	r:=rect{width:3,height:4}
	c:=circle{radius:5}
	//rect和circle均实现了geometry接口，所以可以作为measure的参数
	measure(r)
	measure(c)
}