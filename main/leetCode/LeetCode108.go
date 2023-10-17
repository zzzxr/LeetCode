package main

/*
将有序数组转换为二叉搜索树
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
*/
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}

func sortedArrayToBSTHelper(nums []int, i, j int) *TreeNode {
	if i > j {
		return nil
	}
	mid := (i + j) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = sortedArrayToBSTHelper(nums, i, mid-1)
	root.Right = sortedArrayToBSTHelper(nums, mid+1, j)
	return root
}
