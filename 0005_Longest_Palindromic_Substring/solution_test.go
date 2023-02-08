package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "Example 1",
			s:        "babad",
			expected: "bab",
		},
		{
			name:     "Example 2",
			s:        "cbbd",
			expected: "bb",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			m := longestPalindrome(test.s)

			g.Expect(m).
				Should(Equal(test.expected))
		})
	}
}
