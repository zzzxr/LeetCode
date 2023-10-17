package main

import "fmt"

/*
需求：测试协程和管道的使用
（1）用协程writeData() 向管道 intChan 写入50个数
（2）用协程readData() 从管道 intChan 中读取这50个数
*/

func writeData(ch1 chan int) {
	for i := 0; i < 50; i++ {
		go func(i int) {
			ch1 <- i
		}(i)
	}
}

func readData(ch1 chan int, ch2 chan bool) {
	num := <-ch1
	fmt.Println("num = ", num)
	close(ch1)
	ch2 <- true
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	if status := <-exitChan; status {
		fmt.Println("end")
	}
}
