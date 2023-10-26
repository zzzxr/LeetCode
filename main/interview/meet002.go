package main

import (
	"fmt"
)

func Solution1(N int) {
	var enable_print int
	enable_print = N % 10
	for N > 0 {
		if enable_print == 0 && N/10 != 0 {
			enable_print = 1
		} else {
			fmt.Print(N % 10)
		}
		N = N / 10
	}
}
func main() {
	Solution1(10000)
}
