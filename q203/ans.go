/**
203
*/

package q203

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var _ = new(utils.Util)

type ListNode = utils.ListNode

func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

// 2ms 4.7mb
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return head
}

func Solve(head *ListNode, val int) *ListNode {
	return removeElements(head, val)
}
