package main

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var result bytes.Buffer
	for row := 0; row < numRows; row++ {
		i := row
		j := 0 // odd for vertical, even for diagonal
		for i < len(s) {
			if j%2 == 0 {
				result.WriteByte(s[i])
			} else {
				i += (numRows - 1) * 2
				if row != 0 && row != numRows-1 && i-row*2 < len(s) {
					result.WriteByte(s[i-row*2])
				}
			}
			j++
		}
	}
	return result.String()
}
