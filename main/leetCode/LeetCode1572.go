package main

/*
矩阵对角线元素的和
给你一个正方形矩阵 mat，请你返回矩阵对角线元素的和。

请你返回在矩阵主对角线上的元素和副对角线上且不在主对角线上元素的和。
输入：mat = [[1,2,3],

	[4,5,6],
	[7,8,9]]

输出：25
解释：对角线的和为：1 + 5 + 9 + 3 + 7 = 25
请注意，元素 mat[1][1] = 5 只会被计算一次。
*/
func diagonalSum(mat [][]int) int {
	m := len(mat)
	sum := 0
	for i := 0; i < m; i++ {
		if i == m-1-i {
			sum += mat[i][i]
		} else {
			sum += mat[i][i] + mat[i][m-1-i]
		}
	}
	return sum
}
