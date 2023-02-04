package main

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	resultNext := result
	addendum := 0

	for addendum > 0 || l1 != nil || l2 != nil {
		sum := addendum

		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		if sum > 0 || l1 != nil || l2 != nil {
			resultNext.Next = &ListNode{
				Val: sum % 10,
			}
			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
			resultNext = resultNext.Next
			addendum = sum / 10
		}
	}

	return result.Next
}
