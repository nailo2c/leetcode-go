package problem0002

// ListNode is a struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dfs func(*ListNode, *ListNode, int) *ListNode
	dfs = func(l1 *ListNode, l2 *ListNode, k int) *ListNode {
		if l1 == nil && l2 == nil {
			if k == 1 {
				return &ListNode{Val: 1, Next: nil}
			}
			return nil
		}

		sum := k
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum >= 10 {
			sum -= 10
			k = 1
		} else {
			k = 0
		}
		head := &ListNode{Val: sum}
		head.Next = dfs(l1, l2, k)

		return head
	}

	return dfs(l1, l2, 0)
}
