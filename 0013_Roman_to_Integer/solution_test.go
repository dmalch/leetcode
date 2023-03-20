package main

import (
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
			name:     "Example 1",
			s:        "III",
			expected: 3,
		},
		{
			name:     "Example 2",
			s:        "LVIII",
			expected: 58,
		},
		{
			name:     "Example 3",
			s:        "MCMXCIV",
			expected: 1994,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			s := romanToInt(test.s)

			g.Expect(s).
				Should(Equal(test.expected))
		})
	}
}
