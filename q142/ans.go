/**
142
*/

package q142

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

type ListNode = utils.ListNode

func detectCycle(head *ListNode) *ListNode {
	var (
		fast = head
		slow = head
	)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			var p = head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
}
