package main

import "fmt"

func main() {
	temp := map[float64]float64{3.00000000000000001: 3.14, 3.00000000000000002: 3.14}

	for _, i := range temp {
		fmt.Println(i)
	}
}
