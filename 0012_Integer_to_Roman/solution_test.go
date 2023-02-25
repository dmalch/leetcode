package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		i        int
		expected string
	}{
		{
			name:     "Example 1",
			i:        3,
			expected: "III",
		},
		{
			name:     "Example 2",
			i:        58,
			expected: "LVIII",
		},
		{
			name:     "Example 2",
			i:        1994,
			expected: "MCMXCIV",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			s := intToRoman(test.i)

			g.Expect(s).
				Should(Equal(test.expected))
		})
	}
}
