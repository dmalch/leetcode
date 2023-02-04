package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test1(t *testing.T) {
	g := NewGomegaWithT(t)

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{Val: 3},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{Val: 4},
		},
	}

	sum := addTwoNumbers(l1, l2)

	g.Expect(sum).
		Should(Equal(
			&ListNode{
				Val: 7,
				Next: &ListNode{
					Val:  0,
					Next: &ListNode{Val: 8},
				},
			},
		))
}

func Test2(t *testing.T) {
	g := NewGomegaWithT(t)

	l1 := &ListNode{
		Val: 0,
	}

	l2 := &ListNode{
		Val: 0,
	}

	sum := addTwoNumbers(l1, l2)

	g.Expect(sum).
		Should(Equal(
			&ListNode{
				Val: 0,
			},
		))
}

func Test3(t *testing.T) {
	g := NewGomegaWithT(t)

	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}

	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}

	sum := addTwoNumbers(l1, l2)

	g.Expect(sum).
		Should(Equal(
			&ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 1,
										},
									},
								},
							},
						},
					},
				},
			},
		))
}
