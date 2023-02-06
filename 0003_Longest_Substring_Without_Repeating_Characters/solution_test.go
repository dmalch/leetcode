package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	tests := []struct {
		name           string
		str            string
		expectedLength int
	}{
		{
			name:           "Example 1",
			str:            "abcabcbb",
			expectedLength: 3,
		},
		{
			name:           "Example 2",
			str:            "bbbbb",
			expectedLength: 1,
		},
		{
			name:           "Example 3",
			str:            "pwwkew",
			expectedLength: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			l := lengthOfLongestSubstring(test.str)

			g.Expect(l).
				Should(Equal(test.expectedLength))
		})
	}
}
