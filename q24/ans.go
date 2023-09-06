/**
24
*/

package q24

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

type ListNode = utils.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var cur, n, nn *ListNode
	cur = head
	n = head.Next
	nn = n.Next
	n.Next = cur
	cur.Next = swapPairs(nn)
	return n
}

func swapPairs_2(head *ListNode) *ListNode {
	/*if head == nil || head.Next == nil {
		return head
	}*/
	res := &ListNode{Val: 0, Next: head}
	cur := res
	for cur.Next != nil && cur.Next.Next != nil {
		n1 := cur.Next
		n2 := cur.Next.Next
		n3 := cur.Next.Next.Next
		n1.Next = n3
		cur.Next = n2
		cur.Next.Next = n1
		cur = cur.Next.Next
	}
	return res.Next
}

func Solve(in []int) []int {
	return util.LinkedListToSlice(swapPairs(util.SliceToLinkedList(in)))
}
