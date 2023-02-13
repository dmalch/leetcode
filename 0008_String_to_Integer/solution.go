package main

import "math"

func myAtoi(s string) int {
	if s == "" {
		return 0
	}

	result := 0

	i := 0

	// skip leading white spaces
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}

	if i >= len(s) {
		return 0
	}

	sign := 1
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}

		num := int(s[i] - '0')
		num *= sign

		result *= 10

		if sign > 0 && math.MaxInt32-result < num {
			result = math.MaxInt32
			break
		} else if math.MinInt32-result > num {
			result = math.MinInt32
			break
		}
		result += num
	}

	return result
}
