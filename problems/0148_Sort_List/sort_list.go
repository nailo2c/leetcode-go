package problem0148

// ListNode is a struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	rightHead := slow.Next
	slow.Next = nil

	return merge(sortList(head), sortList(rightHead))
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	prevNode := &ListNode{Val: 0}
	result := prevNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			result.Next = list1
			list1 = list1.Next
		} else {
			result.Next = list2
			list2 = list2.Next
		}
		result = result.Next
	}

	if list1 == nil {
		result.Next = list2
	}

	if list2 == nil {
		result.Next = list1
	}

	return prevNode.Next
}
