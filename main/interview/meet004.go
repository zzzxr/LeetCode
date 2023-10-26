package main

import "fmt"

func Solution4(A []int) int {
	// 创建一个映射来存储累积和的出现次数
	sumMap := make(map[int]int)
	// 初始化累积和为0，以处理从0开始的片段
	sum := 0
	sumMap[0] = 1
	count := 0
	for _, num := range A {
		// 计算累积和
		sum += num

		// 如果当前累积和已经在之前出现过，那么说明在这之间的子数组和为0
		if value, exists := sumMap[sum]; exists {
			count += value
		}

		// 更新累积和出现的次数
		sumMap[sum]++
		// 如果累积和次数过多，返回-1
		if count > 1000000000 {
			return -1
		}
	}

	return count
}

func main() {
	A := []int{2, -2, 0, 3, 4, -7}
	fmt.Println(Solution4(A))
}
