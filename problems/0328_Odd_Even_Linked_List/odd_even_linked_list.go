package problem0328

// ListNode is a struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
