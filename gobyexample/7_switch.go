package gobyexample
import "fmt"
import "time"
func main(){
	i:=2
	fmt.Print("write",i," as ")
	switch i{
	case 1: fmt.Println("one")
	case 2: fmt.Println("two")
	case 3: fmt.Println("three")
	}
	//可以在同一个case语句中使用逗号来分隔多个表达式
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("it's the weekend")
		//这里也使用了default case
	default:
		fmt.Println("it's a weekday")
	}
	t:=time.Now()
	//没有表达式的switch是另一种实现if/else逻辑的路子，同时这也展示了case表达式可以是非常量值。
	switch{
	case t.Hour()<12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
	//类型switch比较了类型而非值
	whatAmI := func(i interface{}){
		switch t:=i.(type){
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type %T\n",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}