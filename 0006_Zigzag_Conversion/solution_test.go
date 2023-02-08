package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		numRows  int
		expected string
	}{
		{
			name:     "Example 1",
			s:        "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "Example 2",
			s:        "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			name:     "Example 3",
			s:        "A",
			numRows:  1,
			expected: "A",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			s := convert(test.s, test.numRows)

			g.Expect(s).
				Should(Equal(test.expected))
		})
	}
}
