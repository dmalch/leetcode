package main

import "math"

func reverse(x int) int {
	var res int

	for x != 0 {
		num := x % 10
		x /= 10

		res *= 10

		if math.MaxInt32-res < num || res-math.MinInt32 < num {
			return 0
		}

		res += num
	}

	return res
}
