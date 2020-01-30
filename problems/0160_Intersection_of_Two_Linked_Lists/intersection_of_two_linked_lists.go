package problem0160

// ListNode is a struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	hasLinkedToA, hasLinkedToB := false, false
	for a != nil && b != nil {
		if a == b {
			return a
		}
		a, b = a.Next, b.Next
		if a == nil && !hasLinkedToB {
			a = headB
			hasLinkedToB = true
		}
		if b == nil && !hasLinkedToA {
			b = headA
			hasLinkedToA = true
		}
	}
	return nil
}
