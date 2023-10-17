package main

/*
x 的平方根
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
*/
func mySqrt(x int) int {
	r := x
	l := 0
	ans := -1
	for l < r {
		mid := l + (l-r)/2
		if mid*mid <= x {
			ans = mid
			l += 1
		} else {
			r -= 1
		}
	}
	return ans
}
