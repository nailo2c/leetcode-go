package problem0141

// ListNode is a struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	}

	slow, fast := head, head.Next.Next

	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}

	return false
}
