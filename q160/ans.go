/**
160
*/

package q160

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

type ListNode = utils.ListNode

func init() {
	util = utils.NewUtil()
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var la, lb int
	var (
		pa = headA
		pb = headB
	)
	for pa != nil {
		la++
		pa = pa.Next
	}
	for pb != nil {
		lb++
		pb = pb.Next
	}
	pa = headA
	pb = headB
	if la > lb {
		for i := 0; i < la-lb; i++ {
			pa = pa.Next
		}
	} else {
		for i := 0; i < lb-la; i++ {
			pb = pb.Next
		}
	}
	for pa != nil {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
	}
	return nil
}

func Solve() any {
	return nil
}
