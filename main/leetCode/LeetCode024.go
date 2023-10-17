package main

/*
24. 两两交换链表中的节点
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
输入：head = [1,2,3,4]
输出：[2,1,4,3]
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	tempList := dummyHead
	for tempList.Next != nil && tempList.Next.Next != nil {
		node1 := tempList.Next
		node2 := tempList.Next.Next
		tempList.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		tempList = node2.Next
	}
	return dummyHead.Next
}
