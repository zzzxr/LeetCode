package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
面试题
在10万个数字中寻找前100大 (top 100)的数字，要求在20万次比较中得到答案
*/

var
func topKFrequent(nums []int, k int) []int {
	m := 0 // 比较次数
	if len(nums) < k {
		return nums
	}
	if len(nums) == 0 {
		return nil
	}
	// 构建一个二叉树
	root := buildTreeNode(nums)
}

func buildTreeNode(nums []int) *TreeNode {
	root := &TreeNode{}
	for i := 0; i < len(nums); i++ {
		root.Val = nums[i]
	}
}

func main() {
	// 生成随机数
	rand.Seed(time.Now().UnixNano())
	array := make([]int, 100000)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(10000)
	}
	fmt.Println(topKFrequent(array, 100))
}
