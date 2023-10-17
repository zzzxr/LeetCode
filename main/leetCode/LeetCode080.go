package main

import "fmt"

/*
删除有序数组中的重复项 II
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 不需要考虑数组中超出新长度后面的元素。
*/
func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	i, j := 2, 2
	for i < len(nums) {
		if nums[i-2] == nums[j] {
			nums[i] = nums[j]
			j++
		}
		i++

	}
	return j
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates2(nums))
}
