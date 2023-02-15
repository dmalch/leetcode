package main

import (
	"math"
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected bool
	}{
		{
			name:     "Example 1",
			x:        121,
			expected: true,
		},
		{
			name:     "Example 2",
			x:        -121,
			expected: false,
		},
		{
			name:     "Example 3",
			x:        10,
			expected: false,
		},
		{
			name:     "Example 4",
			x:        0,
			expected: true,
		},
		{
			name:     "Example 5",
			x:        1221,
			expected: true,
		},
		{
			name:     "Example 6",
			x:        math.MaxInt32,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			s := isPalindrome(test.x)

			g.Expect(s).
				Should(Equal(test.expected))
		})
	}
}
