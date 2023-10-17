package main

import (
	"fmt"
	"math"
)

/*
下降路径最小和 II
给你一个 n x n 整数矩阵 grid ，请你返回 非零偏移下降路径 数字和的最小值。

非零偏移下降路径 定义为：从 grid 数组中的每一行选择一个数字，且按顺序选出来的数字中，相邻数字不在原数组的同一列。
输入：grid = [[1,2,3],[4,5,6],[7,8,9]]
输出：13
解释：
所有非零偏移下降路径包括：
[1,5,9], [1,5,7], [1,6,7], [1,6,8],
[2,4,8], [2,4,9], [2,6,7], [2,6,8],
[3,4,8], [3,4,9], [3,5,7], [3,5,9]
下降路径中数字和最小的是 [1,5,7] ，所以答案是 13 。
*/

func main() {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(minFallingPathSum(grid))
}
func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 0)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	for i := 0; i < n; i++ {
		dp[0][i] = grid[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if j == k {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k]+grid[i][j])
			}
		}
	}
	result := math.MaxInt
	for i := 0; i < n; i++ {
		result = min(result, dp[n-1][i])
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
