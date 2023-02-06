package main

func lengthOfLongestSubstring(s string) int {
	pos1 := 0
	pos2 := 0

	maxLength := 0

	chars := make(map[byte]int)
	for pos2 < len(s) {
		pos2++

		charAt := s[pos2-1]

		if _, ok := chars[charAt]; ok {
			chars[charAt]++
			for chars[charAt] > 1 {
				chars[s[pos1]]--
				pos1++
			}
		} else {
			chars[charAt] = 1
		}

		length := pos2 - pos1
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
