package main

import (
	"fmt"
)

/*
实现生产者和消费者
*/

func producer(ch chan int) {
	for i := 0; i < 100; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}
}

func consumer(ch chan int) {
	for {
		select {
		case num := <-ch:
			fmt.Println("消费数字为：", num)
			wg.Done()
		}
	}
}

func main() {
	ch := make(chan int, 100)
	wg.Add(100)
	go producer(ch)
	go consumer(ch)
	wg.Wait()
}
