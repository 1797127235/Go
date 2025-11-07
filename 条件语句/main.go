package main

import "fmt"

func add(x int,y int) int {
	return x + y	
}

func main() {
	var num int = 10
	if num < 0 {
		fmt.Println("num is less than 0")
	} else if num == 0 {
		fmt.Println("num is equal to 0")
	} else {
		fmt.Println("num is greater than 0")
	}

	for i :=0 ; i < 10 ; i++ {
		fmt.Println(i)
	}

}
