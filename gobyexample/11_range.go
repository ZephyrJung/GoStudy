package gobyexample
import "fmt"
func main(){
	//这里使用range来计算元素的和，数组也是类似的用法
	nums:=[]int{2,3,4}
	sum:=0
	for _,num:=range nums{
		sum+=num
	}
	fmt.Println("sum:",sum)
	//在slice和array上的range均为每个条目提供了索引和值
	for i,num:=range nums{
		if num==3{
			fmt.Println("index",i)
		}
	}
	//map上的range通过key/value对进行遍历
	kvs:=map[string]string{"a":"apple","b":"banana"}
	for k,v:=range kvs{
		fmt.Printf("%s -> %s\n",k,v)
	}
	//range可以仅通过key进行遍历
	for k:= range kvs{
		fmt.Println("key:",k)
	}
	//range作用在string上将得到unicode code points，第一个值是字符的起始字节索引，第二个值是字符本身
	for i,c:= range "go" {
		fmt.Println(i,c)
	}
}