package main

import (
	"math"
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 0",
			s:        "",
			expected: 0,
		},
		{
			name:     "Example 0.1",
			s:        " ",
			expected: 0,
		},
		{
			name:     "Example 1",
			s:        "42",
			expected: 42,
		},
		{
			name:     "Example 1.1",
			s:        "123",
			expected: 123,
		},
		{
			name:     "Example 2",
			s:        "   -42",
			expected: -42,
		},
		{
			name:     "Example 2.1",
			s:        "+-12",
			expected: 0,
		},
		{
			name:     "Example 3",
			s:        "4193 with words",
			expected: 4193,
		},
		{
			name:     "Example 4",
			s:        "21474836460",
			expected: math.MaxInt32,
		},
		{
			name:     "Example 5",
			s:        "-2147483649",
			expected: math.MinInt32,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			s := myAtoi(test.s)

			g.Expect(s).
				Should(Equal(test.expected))
		})
	}
}
