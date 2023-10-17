package main

import "fmt"

func Solution(A []int) int {
	// Implement your solution here
	temp := 1
	numMap := make(map[int]struct{})
	for _, a := range A {
		numMap[a] = struct{}{}
	}
	for {
		if _, ok := numMap[temp]; !ok {
			break
		}
		temp++
	}
	return temp
}

func main() {
	A := []int{1, 3, 6, 4, 1, 2}
	fmt.Println(Solution(A))
}
