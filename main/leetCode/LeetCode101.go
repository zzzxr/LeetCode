package main

/*
	对称二叉树

给你一个二叉树的根节点 root ， 检查它是否轴对称。
*/
func isSymmetric(root *TreeNode) bool {

	return swapTree(root.Right, root.Left) && swapTree(root.Left, root.Right)
}

func swapTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return swapTree(p.Right, q.Left) && swapTree(q.Left, q.Right)
}
