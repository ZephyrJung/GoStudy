package gobyexample
import "fmt"
func fact(n int) int {
	if n==0 {
		return 1
	}
	//fact函数调用自身，直到n为0
	return n * fact(n-1)
}
func main(){
	fmt.Println(fact(7))
}