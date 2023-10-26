package main

import "fmt"

/*
LeetCode 1615. 最大网络秩
*/
func Solution3(A []int, B []int, N int) int {
	// 创建一个城市对应的排名映射
	ranks := make(map[int]int)

	// 遍历道路，计算城市的排名
	for i := 0; i < len(A); i++ {
		ranks[A[i]]++
		ranks[B[i]]++
	}

	maxRank := 0

	// 找到最大的排名
	for _, rank := range ranks {
		if rank > maxRank {
			maxRank = rank
		}
	}

	return maxRank
}

func main() {
	A := []int{1, 2, 3, 3}
	B := []int{2, 3, 1, 4}
	fmt.Println(Solution3(A, B, 4))
}
