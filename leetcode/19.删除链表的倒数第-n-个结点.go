/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var res = head
	i, j := head, head
	for n > 0 {
		j = j.Next
		n--
	}
	if j == nil {
		return head.Next
	}
	for j.Next != nil {
		i = i.Next
		j = j.Next
	}
	i.Next = i.Next.Next
	return res
}

// @lc code=end

