/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	tmp := 0
	for l1 != nil {
		a := (l1.Val + l2.Val + tmp) % 10
		b := (l1.Val + l2.Val + tmp) / 10
		tmp = b
		l1.Val = a
		if l2.Next != nil && l1.Next == nil {
			l1.Next = &ListNode{}
		}
		if l1.Next != nil && l2.Next == nil {
			l2.Next = &ListNode{}
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if tmp == 1 {
		l1 = head
		for l1.Next != nil {
			l1 = l1.Next
		}
		l1.Next = &ListNode{Val: 1}
	}
	return head
}

// @lc code=end

