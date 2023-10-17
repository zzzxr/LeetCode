package main

import (
	"fmt"
	"sync"
)

/*
开协程打印，计数器控制程序结束
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
