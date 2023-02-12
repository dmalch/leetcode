package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
		s    string
	}{
		{
			name: "Example 1",
		},
		{
			name: "Example 2",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			s := myAtoi(test.s)

			g.Expect(s).
				Should(Equal(0))
		})
	}
}
