/**
203
*/

package q203

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var _ = new(utils.Util)

type ListNode = utils.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

func Solve(head *ListNode, val int) *ListNode {
	return removeElements(head, val)
}
