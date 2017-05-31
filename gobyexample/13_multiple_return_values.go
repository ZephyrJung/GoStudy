package gobyexample
import "fmt"
//方法签名中的(int,int)表明它将返回两个int值
func vals()(int,int){
	return 3,7
}
func main(){
	//这里我们通过多重赋值来使用两个不同的返回值
	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)
	//如果只需要返回结果的子集，使用空白占位符_
	_,c := vals()
	fmt.Println(c)
}
