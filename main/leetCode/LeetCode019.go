package main

/*
删除链表的倒数第 N 个结点
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	p := dummy
	for i := 0; i < length-n; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return p
}

func getLength(head *ListNode) int {
	num := 0
	for ; head.Next != nil; head = head.Next {
		num++
	}
	return num
}
