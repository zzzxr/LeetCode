package main

import (
	"fmt"
	"time"
)

/*
map默认是并发不安全的，同时对map进行并发读写时，程序会panic，原因如下：

Go 官方在经过了长时间的讨论后，认为 Go map 更应适配典型使用场景（不需要从多个 goroutine 中进行安全访问），而不是为了小部分情况（并发访问），导致大部分程序付出加锁代价（性能），决定了不支持。

场景: 2个协程同时读和写，以下程序会出现致命错误：fatal error: concurrent map writes
*/

func main() {
	tempMap := map[int]int{}
	for i := 0; i < 100; i++ {
		go func(i int) {
			tempMap[i] = i
		}(i)
	}
	for j := 0; j < 100; j++ {
		go func(j int) {
			fmt.Println(tempMap[j])
		}(j)
	}
	time.Sleep(time.Second)
}
