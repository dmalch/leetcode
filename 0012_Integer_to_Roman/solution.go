package main

import "strings"

func intToRoman(num int) string {
	var result strings.Builder

	// get thousands
	m := num / 1000
	for i := 0; i < m; i++ {
		result.WriteString("M")
	}

	num %= 1000

	// get hundreds
	m = num / 100
	if m == 9 {
		result.WriteString("CM")
	} else if m == 4 {
		result.WriteString("CD")
	} else if m >= 5 && m < 9 {
		result.WriteString("D")
		for i := 5; i < m; i++ {
			result.WriteString("C")
		}
	} else if m >= 1 && m < 4 {
		for i := 0; i < m; i++ {
			result.WriteString("C")
		}
	}

	num %= 100

	// get decimals
	m = num / 10
	if m == 9 {
		result.WriteString("XC")
	} else if m == 4 {
		result.WriteString("XL")
	} else if m >= 5 && m < 9 {
		result.WriteString("L")
		for i := 5; i < m; i++ {
			result.WriteString("X")
		}
	} else if m >= 1 && m < 4 {
		for i := 0; i < m; i++ {
			result.WriteString("X")
		}
	}

	num %= 10

	// get decimals
	if num == 9 {
		result.WriteString("IX")
	} else if num == 4 {
		result.WriteString("IV")
	} else if num >= 5 && num < 9 {
		result.WriteString("V")
		for i := 5; i < num; i++ {
			result.WriteString("I")
		}
	} else if num >= 1 && num < 4 {
		for i := 0; i < num; i++ {
			result.WriteString("I")
		}
	}

	return result.String()
}
