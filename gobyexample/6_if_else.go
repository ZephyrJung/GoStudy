package gobyexample
import "fmt"
func main(){
	if 7%2==0{
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}
	//可以单独使用if
	if 8%4==0{
		fmt.Println("8 is divisible by 4")
	}
	//条件之前可以有语句，声明在该语句中的任何变量可以用于所有分支
	if num:=9;num<0{
		fmt.Println(num,"is negative")
	}else if num<10{
		fmt.Println(num,"has 1 digit")
	}else{
		fmt.Println(num,"has multiple digits")
	}
}