package main

import "fmt"

/*
map中get的两种写法
*/
func main() {
	ageMap := make(map[string]int)
	ageMap["qcrao"] = 18

	// 不带 comma 用法
	age1 := ageMap["stefno"]
	fmt.Println(age1)

	// 带 comma 用法
	age2, ok := ageMap["stefno"]
	fmt.Println(age2, ok)
}
