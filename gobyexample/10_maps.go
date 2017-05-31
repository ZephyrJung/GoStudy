package gobyexample
import "fmt"
func main(){
	//要创建一个空的map，使用内置make函数：make(map[key-type]val-type)
	m:=make(map[string]int)
	//通过经典的name[key]=val语法来设置key/value对
	m["k1"]=7
	m["k2"]=13
	fmt.Println("map:",m)
	//通过name[key]来获取一个value
	v1:=m["k1"]
	fmt.Println("v1: ",v1)
	//len函数返回map中键值对的个数
	fmt.Println("len:",len(m))
	//内置的delete函数将移除map中的键值对
	delete(m,"k2")
	fmt.Println("map:",m)
	//第一个值是该key的value，但此处不需要，故使用空白占位符“_”忽略
	//第二个可选的返回值表明该键是否在map中，这样可以消除不存在的键，和键值为0或者""的歧义
	_,prs:=m["k2"]
	fmt.Println("prs:",prs)
	//声明并初始化map
	n:=map[string]int{"foo":1,"bar":2}
	fmt.Println("map:",n)
}