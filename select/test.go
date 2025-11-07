package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{})

func event() {
	fmt.Println("执行开始")
	time.Sleep(2 * time.Second)
	fmt.Println("执行结束")
	//close(done)

	done <- struct{}{}
}

func main() {
	// go event()
	// select {
	// 	case <-done:
	// 		fmt.Println("事件结束")
	// 	case <-time.After(3 * time.Second):
	// 		fmt.Println("事件超时")
	// }

	ch := make(chan int,1)
	select{
	case ch <- 10:
		fmt.Println("发送成功")
	case val := <- ch:
		fmt.Println("接收成功",val)
	default:
		fmt.Println("默认")
	}
}