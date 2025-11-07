package main
/*
	接口变量包括两个部分
	1.动态类型
	2.动态值
*/
import "fmt"

type Sing interface{
	Sing() 	
}

type Chicken struct{
	Name string
}

func(c Chicken)Sing(){
	fmt.Println(c.Name)
}

func sing(s Sing){
	s.Sing()
}

func print(x interface{}){
	fmt.Printf("类型: %T, 值: %v\n",x,x)
}


func main(){

	ch := Chicken{Name: "c"}
	bh := Chicken{Name: "b"}

	sing(ch)
	sing(bh)

	print(ch)
}