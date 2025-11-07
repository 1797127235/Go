package main

import "fmt"

func main() {
	// var num1 int = 10
	// num2 := 20
	// fmt.Println(num1)
	// fmt.Println(num2)

	// const LENGTH int = 10
	// const WIDTH int = 5
	// var area int
	// const a,b,c = 1,false,"str"

	// area = LENGTH * WIDTH
	// fmt.Printf("面积为 : %d\n",area)

	const (
		a = iota
		b
		c
	)

	fmt.Println(a,b,c)
}

