package gobyexample
import "fmt"
func main(){
	//这里创建了一个5个int元素的数组。默认是零值，对于int而言就是0
	var a [5]int
	fmt.Println("emp:",a)
	//可以通过array[index]=value语法设值，或者通过array[index]取值
	a[4]=100
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])
	//内置的len函数返回数组的长度
	fmt.Println("len:",len(a))
	//声明并初始化数组
	b:=[5]int{1,2,3,4,5}
	fmt.Println("dc1:",b)
	//数组是一维的，但你可以组合类型来构建多维数组结构
	var twoD [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d: ",twoD)
}