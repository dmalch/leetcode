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
		expected int
	}{
		{
			name:     "Example 1",
			x:        123,
			expected: 321,
		},
		{
			name:     "Example 2",
			x:        -123,
			expected: -321,
		},
		{
			name:     "Example 3",
			x:        120,
			expected: 21,
		},
		{
			name:     "Example 4",
			x:        math.MaxInt32,
			expected: 0,
		},
		{
			name:     "Example 5",
			x:        math.MinInt32,
			expected: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			res := reverse(test.x)

			g.Expect(res).
				Should(Equal(test.expected))
		})
	}
}
