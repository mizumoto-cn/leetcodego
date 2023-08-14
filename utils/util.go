package utils

type Util interface {
	SliceToLinkedList(slice []int) *ListNode
	LinkedListToSlice(head *ListNode) []int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type someU struct{}

func NewUtil() Util {
	return &someU{}
}

func (u *someU) SliceToLinkedList(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}
	head := &ListNode{Val: slice[0]}
	cur := head
	for i := 1; i < len(slice); i++ {
		cur.Next = &ListNode{Val: slice[i]}
		cur = cur.Next
	}
	return head
}

func (u *someU) LinkedListToSlice(head *ListNode) []int {
	slice := make([]int, 0)
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	return slice
}
