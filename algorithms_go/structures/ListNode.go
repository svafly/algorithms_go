package structures

import (
	"fmt"
)

//链表节点
type ListNode struct {
	val  int
	Next *ListNode
}

//链表转数组
func Lists2Ints(head *ListNode) []int {
	limit := 100
	counter := 0
	res := []int{}
	for head != nil {
		counter++
		if counter >= limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}
		res = append(res, head.val)
		head = head.Next
	}
	return res
}

//数组转链表
func Ints2Lists(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{val: v}
		t = t.Next
	}
	return l.Next
}
