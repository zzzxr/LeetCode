package main

import (
	"fmt"
	"sync"
)

func m(list []int, f func(x int) (int, error)) ([]int, error) {
	var ch = make(chan int, 10)
	var wg sync.WaitGroup
	rList := make([]int, 0)
	for i, j := range list {
		wg.Add(1)
		go func(i, j int) {
			defer wg.Done()
			ch <- i
			val, err := f(j)
			if err != nil {
				return
			}
			rList = append(rList, val)
			<-ch
		}(i, j)
	}
	wg.Wait()
	return rList, nil
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(m(list, func(x int) (int, error) {
		return x + 1, nil
	}))
}
