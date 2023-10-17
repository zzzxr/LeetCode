package main

import (
	"fmt"
	"sync"
)

/*
交替打印奇数和偶数
方法一：使用无缓冲的channel进行协程间通信
方法二：使用有缓冲的channel
N个协程打印1到maxVal
交替打印字符和数字
交替打印字符串
方法一使用无缓冲的channel
三个协程打印ABC
Channel练习
*/
func printNumber() {
	var wg sync.WaitGroup
	ch := make(chan struct{})
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch <- struct{}{}
			if i%2 == 1 {
				fmt.Println(i, "为奇数")
			}
		}

	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println(i, "为偶数")
			}
		}
	}()
	wg.Wait()
}

func main() {
	printNumber()
}
