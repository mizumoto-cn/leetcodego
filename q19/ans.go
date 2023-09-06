/**
19
*/

package q19

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

type ListNode = util.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {

}

func Solve(in []int, n int) []int {
	return util.LinkedListToSlice(removeNthFromEnd(util.SliceToLinkedList(in), n))
}
