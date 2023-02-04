package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test1(t *testing.T) {
	g := NewGomegaWithT(t)

	sum := twoSum([]int{2, 7, 11, 15}, 9)

	g.Expect(sum).Should(ConsistOf(0, 1))
}

func Test2(t *testing.T) {
	g := NewGomegaWithT(t)

	sum := twoSum([]int{3, 2, 4}, 6)

	g.Expect(sum).Should(ConsistOf(1, 2))
}

func Test3(t *testing.T) {
	g := NewGomegaWithT(t)

	sum := twoSum([]int{3, 3}, 6)

	g.Expect(sum).Should(ConsistOf(0, 1))
}
