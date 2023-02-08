package main

func longestPalindrome(s string) string {
	maxSubstr := ""

	for i := range s {
		// shift right until the letter is the same
		rightPointer := i
		for rightPointer < len(s) && s[rightPointer] == s[i] {
			rightPointer++
		}

		// extend the substring while the letter is the same
		leftPointer := i - 1
		for leftPointer >= 0 && rightPointer < len(s) && s[leftPointer] == s[rightPointer] {
			leftPointer--
			rightPointer++
		}

		leftPointer++

		length := rightPointer - leftPointer
		if length > len(maxSubstr) {
			maxSubstr = s[leftPointer:rightPointer]
		}
	}

	return maxSubstr
}
