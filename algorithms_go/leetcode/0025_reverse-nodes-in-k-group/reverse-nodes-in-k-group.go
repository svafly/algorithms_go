package leetcode

import "algorithms_go/structures"

type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	//边界判定
	if head == nil || head.Next == nil {
		return head
	}
	//取k长的链表
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}
	//反转链表
	newHead := reverseK(head, tail)
	//子链表之间的连接
	head.Next = reverseKGroup(tail, k)
	return newHead
}

//链表反转，这里的尾节点本来是空，现在直接用下一个链表段的头来代替
func reverseK(head *ListNode, tail *ListNode) *ListNode {
	pre := tail
	for head != tail {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
