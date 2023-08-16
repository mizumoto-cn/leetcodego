/**
206
*/

package q206

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var u = utils.NewUtil()

type ListNode = utils.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func Solve(in []int) []int {
	return u.LinkedListToSlice(reverseList(u.SliceToLinkedList(in)))
}
