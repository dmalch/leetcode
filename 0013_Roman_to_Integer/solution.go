package main

func romanToInt(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch == 'M' {
			res += 1000
		} else if ch == 'D' {
			res += 500
		} else if ch == 'C' {
			if i+1 < len(s) && (s[i+1] == 'M' || s[i+1] == 'D') {
				res -= 100
			} else {
				res += 100
			}
		} else if ch == 'L' {
			res += 50
		} else if ch == 'X' {
			if i+1 < len(s) && (s[i+1] == 'C' || s[i+1] == 'L') {
				res -= 10
			} else {
				res += 10
			}
		} else if ch == 'V' {
			res += 5
		} else if ch == 'I' {
			if i+1 < len(s) && (s[i+1] == 'X' || s[i+1] == 'V') {
				res -= 1
			} else {
				res += 1
			}
		}
	}

	return res
}
