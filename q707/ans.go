/**
707
*/

package q707

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var _ = new(utils.Util)

type MyLinkedList struct {
	head *utils.ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.head
	if cur == nil {
		return -1
	}
	for i := 0; i < index; i++ {
		if cur.Next == nil {
			return -1
		}
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.head == nil {
		this.head = &utils.ListNode{Val: val, Next: nil}
		return
	}
	cur := &utils.ListNode{Val: val, Next: this.head}
	this.head = cur
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.head = &utils.ListNode{Val: val, Next: nil}
		return
	}
	cur := this.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &utils.ListNode{Val: val, Next: nil}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	cur := this.head
	for i := 0; i < index-1; i++ {
		if cur == nil || cur.Next == nil {
			return
		}
		cur = cur.Next
	}
	if cur == nil {
		return
	}
	cur.Next = &utils.ListNode{Val: val, Next: cur.Next}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.head = this.head.Next
		return
	}
	cur := this.head
	for i := 0; i < index-1; i++ {
		if cur.Next == nil {
			return
		}
		cur = cur.Next
	}
	if cur == nil || cur.Next == nil {
		return
	}
	cur.Next = cur.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func Solve() bool {
	return true
}

func MyLinkedListToSlice(head *MyLinkedList) []int {
	return utils.NewUtil().LinkedListToSlice(head.head)
}

func SliceToMyLinkedList(s []int) *MyLinkedList {
	node := utils.NewUtil().SliceToLinkedList(s)
	return &MyLinkedList{head: node}
}
