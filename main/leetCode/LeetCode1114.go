package main

import (
	"fmt"
	"time"
)

/*
按序打印
三个不同的线程 A、B、C 将会共用一个 Foo 实例。

线程 A 将会调用 first() 方法
线程 B 将会调用 second() 方法
线程 C 将会调用 third() 方法
请设计修改程序，以确保 second() 方法在 first() 方法之后被执行，third() 方法在 second() 方法之后被执行。
*/
var ch1 = make(chan int)
var ch2 = make(chan int)

func main() {
	go first()
	go second()
	go third()
	time.Sleep(time.Second)

}

func first() {
	fmt.Println("First")
	ch1 <- 1
}
func second() {
	<-ch1
	fmt.Println("Second")
	ch2 <- 1
}
func third() {
	<-ch2
	fmt.Println("Third")
	close(ch2)
}
