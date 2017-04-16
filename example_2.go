package main
import "fmt"
import "math"
const s string="const string"
func main() {
	fmt.Println(s)
	const n=50000000
	const d=3e20/n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
	i:=1
	for i<=3{
		fmt.Println(i)
		i=i+1
	}
	for j:=7;j<=9;j++{
		fmt.Println(j)
	}
	for{
		fmt.Println("loop")
		break
	}
	for n:=0;n<=5;n++{
		if n%2==0{
			continue
		}
		fmt.Println(n)
	}
	if 7%2==0 {
		fmt.Println("7是偶数")
	}else{
		fmt.Pritnln("7是奇数")
	}
	if 8%4==0 {
		fmt.Println("8可以被4正除")
	}
	if num:=9;num<0{
		fmt.Println(num," is negative")
	}else if num<10{
		fmt.Println(num,"has 1 digit")
	}else{
		fmt.Println(num,"has more than 1 digit")
	}
}	