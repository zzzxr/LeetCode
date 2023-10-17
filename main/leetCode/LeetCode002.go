package main

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	var tempNum int
	for l1 != nil || l2 != nil {
		m1, n1 := 0, 0
		if l1 != nil {
			m1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n1 = l2.Val
			l2 = l2.Next
		}

		sum002 := m1 + n1 + tempNum
		tempNum = sum002 / 10
		sum002 = sum002 % 10
		if head == nil {
			head = &ListNode{Val: sum002}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum002}
			tail = tail.Next
		}
	}
	if tempNum > 0 {
		tail.Next = &ListNode{Val: tempNum}
	}

	return head
}
