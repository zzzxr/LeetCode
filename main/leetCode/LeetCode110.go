package main

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(height(root.Right)-height(root.Left)) <= 1 && isBalanced(root.Right) && isBalanced(root.Left) {
		return true
	}
	return false
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Right), height(root.Left)) + 1
}
func abs(x int) int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}
