package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	tests := []struct {
		name string
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

			s := solution()

			g.Expect(s).
				Should(Equal(0))
		})
	}
}
