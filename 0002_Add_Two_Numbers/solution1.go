package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers2(l1, l2, 0)
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode, addendum int) *ListNode {
	sum := addendum
	if l1 != nil {
		sum += l1.Val
	}
	if l2 != nil {
		sum += l2.Val
	}

	if sum > 0 || l1 != nil || l2 != nil {
		var l1Next *ListNode
		if l1 != nil {
			l1Next = l1.Next
		}
		var l2Next *ListNode
		if l2 != nil {
			l2Next = l2.Next
		}

		result := &ListNode{
			Val:  sum % 10,
			Next: addTwoNumbers2(l1Next, l2Next, sum/10),
		}

		return result
	}

	return nil
}
