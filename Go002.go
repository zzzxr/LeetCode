package main

import (
	"fmt"
	"time"
)

/*
有三个协程，编号为123，如何实现顺序打印
*/

func printNumber(number int, ch, nextCh chan struct{}) {
	for {
		<-ch // 等待通道中的信号，以控制打印间隔
		fmt.Printf("goroutine %d\n", number)
		nextCh <- struct{}{}
	}
}

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	// 创建3个goroutine

	go printNumber(1, ch1, ch2)
	go printNumber(2, ch2, ch3)
	go printNumber(3, ch3, ch1)

	ch1 <- struct{}{}
	time.Sleep(time.Second)
}
