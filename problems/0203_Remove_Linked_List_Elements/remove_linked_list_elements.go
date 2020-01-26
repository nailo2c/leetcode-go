package problem0203

// ListNode is a struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	preHead := ListNode{Next: head}

	temp := &preHead
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return preHead.Next
}
