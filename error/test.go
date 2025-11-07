package main

import (
    "errors"
    "fmt"
)

func divida(a,b float64) (float64,error){
	if b == 0{
		return 0,errors.New("除数不能为0")
	}
	return a/b,nil
}

func mayPanic(){
	defer func(){
		if r := recover();r != nil{
			fmt.Println("捕获到 panic:",r)
		}
	}()
	
	fmt.Println("执行中")
	panic("程序崩溃")
	fmt.Println("不会执行")
}

func main() {
	// result,err := divida(1,0)
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)
	mayPanic()

	fmt.Println("main函数结束")
}
