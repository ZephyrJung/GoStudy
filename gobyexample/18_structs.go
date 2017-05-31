package gobyexample
import "fmt"
type person struct{
	name string
	age int
}
func main(){
	//这个语法创建了一个新的struct
	fmt.Println(person{"Bob",20})
	//在初始化struct时，可以指定字段名
	fmt.Println(person{name:"Alice",age:30})
	//被忽略的字段将会被初始化为零
	fmt.Println(person{name:"Fred"})
	//一个&将产生struct的指针
	fmt.Println(&person{name:"Ann",age:40})
	s:=person{name:"Sean",age:50}
	//通过点号访问结构体中的字段
	fmt.Println(s.name)
	sp:=&s
	fmt.Println(sp.age)
	//对于结构体的指针也可以使用点号操作符，指针将会自动解引用
	sp.age=51
	//结构体是可变的
	fmt.Println(sp.age)
}