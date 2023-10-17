package main

import (
	"fmt"
)

/*
var wg sync.WaitGroup
按照顺序打印goroutine123
*/

func main() {

	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)
	wg.Add(3)
	ch1 <- struct{}{}
	go print(ch1, ch2, "gorountine1")
	go print(ch2, ch3, "gorountine2")
	go print(ch3, ch1, "gorountine3")
	wg.Wait()
	fmt.Println("End")
}

func print(input, output chan struct{}, gorouine string) {
	select {
	case <-input:
		fmt.Printf("%s\n", gorouine)
		output <- struct{}{}
	}

	wg.Done()
}
