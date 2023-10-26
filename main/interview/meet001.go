package main

/*
这是一个演示任务。

写一个函数：

func 解决方案(A []int) int

给定一个包含 N 个整数的数组 A，返回 A 中未出现的最小正整数（大于 0）。

例如，给定 A = [1, 3, 6, 4, 1, 2]，该函数应返回 5。

给定 A = [1, 2, 3]，该函数应返回 4。

给定 A = [−1, −3]，该函数应返回 1。

为以下假设编写一个有效的算法：

N是[ 1..100,000 ]范围内的整数；
数组 A 的每个元素都是 [ -1,000,000 .. 1,000,000 ]范围内的整数。
*/

func Solution(A []int) int {
	// Implement your solution here
	intNum := map[int]bool{}
	temp := 1
	for _, num := range A {
		if num < 0 {
			continue
		} else {
			intNum[num] = true
		}
	}
	for {
		if _, ok := intNum[temp]; ok {
			temp++
		} else {
			break
		}
	}
	return temp
}
