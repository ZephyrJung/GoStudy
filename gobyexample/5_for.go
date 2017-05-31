package gobyexample
import "fmt"
func main(){
	i := 1
	//最基本的类型，只有一个条件
	for i<=3{
		fmt.Println(i)
		i = i+1
	}
	//经典的 初始化/条件/循环后 for循环
	for j:=7;j<=9;j++{
		fmt.Println(j)
	}
	//没有条件的for语句将无限循环，除非在内部的方法中使用了break或者return
	for{
		fmt.Println("loop")
		break
	}
	//也可以使用continue来直接到下一个循环
	for n:=0;n<=5;n++ {
		if n%2 == 0{
			continue
		}
		fmt.Println(n)
	}
}