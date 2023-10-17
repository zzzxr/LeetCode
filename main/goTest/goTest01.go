package main

import "fmt"

/*
两个nil一定相等吗
*/

func main() {
	var a *int = nil
	var b *string = nil
	fmt.Println(a == b)
}
