package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	tests := []struct {
		name           string
		nums1          []int
		nums2          []int
		expectedMedian float64
	}{
		{
			name:           "Example 1",
			nums1:          []int{1, 3},
			nums2:          []int{2},
			expectedMedian: 2,
		},
		{
			name:           "Example 2",
			nums1:          []int{1, 2},
			nums2:          []int{3, 4},
			expectedMedian: 2.5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			m := findMedianSortedArrays(test.nums1, test.nums2)

			g.Expect(m).
				Should(Equal(test.expectedMedian))
		})
	}
}
