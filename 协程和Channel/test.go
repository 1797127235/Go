package main

import (
	"fmt"
	"sync"
)

var ch chan int

func shopping(name string, wg *sync.WaitGroup) {
	//defer wg.Done()
	fmt.Printf("%s 开始购物\n", name)
	ch <- 1 // 发送完成信号/金额等
	fmt.Printf("%s 购物结束\n", name)

	wg.Done()
}

func main() {
	ch = make(chan int,3) 

	var wg sync.WaitGroup
	wg.Add(3)

	go shopping("小明", &wg)
	go shopping("小红", &wg)
	go shopping("小刚", &wg)

	// 在所有生产者结束后关闭通道
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 主协程作为消费者收集结果
	monneylist := []int{}
	for v := range ch { // 通道关闭后自动退出循环
		monneylist = append(monneylist, v)
	}

	fmt.Println("结果：", monneylist) // 例如 [1 1 1]
}
