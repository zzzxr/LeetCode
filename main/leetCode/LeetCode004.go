package main

import "fmt"

/*
寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	nums3 := make([]int, m+n)
	p1, p2, p3 := 0, 0, 0
	for p1 < m || p2 < n {
		if p1 == m {
			nums3[p3] = nums2[p2]
			p2++
			p3++
			continue
		}
		if p2 == n {
			nums3[p3] = nums1[p1]
			p1++
			p3++
			continue
		}
		if nums1[p1] <= nums2[p2] {
			nums3[p3] = nums1[p1]
			p1++
			p3++
		} else {
			nums3[p3] = nums2[p2]
			p2++
			p3++
		}
	}

	if (m+n)%2 == 1 {
		return float64(nums3[(m+n)/2])
	} else {
		return float64(nums3[(m+n)/2]+nums3[(m+n)/2-1]) / 2
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
