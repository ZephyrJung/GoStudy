package gobyexample
import "fmt"
func main(){
	//与数组不同，切片仅由其包含的元素（不是元素的数量）键入。
	//要创建一个非零长度的空切片，请使用内置的make。 这里我们制作长度为3的字符串（最初为零值）。
	s:=make([]string,3)
	fmt.Println("emp:",s)
	//可以像数组一样set和get
	s[0]="a"
	s[1]="b"
	s[2]="c"
	fmt.Println("set:",s)
	fmt.Println("get:",s[2])
	//len能够返回slice的长度
	fmt.Println("len:",len(s))
	//作为这些基本操作的补充，slice支持一些令其比数组更丰富的东西。
	//其中一个是append函数，其返回一个包含一个或多个新值的slice
	//注意需要接受append的返回值来获取新的slice值
	s=append(s,"d")
	s=append(s,"e","f")
	fmt.Println("apd:",s)
	//slice可以复制
	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println("cpy:",c)
	//slice支持切片操作，语法为slice[low:high]。如下将得到元素s[2],s[3]和s[4]
	l:=s[2:5]
	fmt.Println("sl1:",l)
	//如下切片截止到（不包含）s[5]
	l=s[:5]
	fmt.Println("sl2:",l)
	//如下切片从s[2]开始（包含）
	l=s[2:]
	fmt.Println("sl3:",l)
	//声明并初始化slice
	t:=[]string{"g","h","i"}
	fmt.Println("dcl:",t)
	//slice也可以组织成一个多为数据结构，内部的slice长度可以变化，这与多维数组不同
	twoD:=make([][]int,3)
	for i:=0;i<3;i++{
		innerLen:=i+1
		twoD[i]=make([]int,innerLen)
		for j:=0;j<innerLen;j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d: ",twoD)
}