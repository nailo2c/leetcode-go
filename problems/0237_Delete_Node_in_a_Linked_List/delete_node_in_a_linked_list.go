package problem0237

// ListNode is a struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	*node = *node.Next
}
