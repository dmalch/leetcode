package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	nums := make([]byte, 0, 1000)

	for x > 0 {
		nums = append(nums, byte(x%10))
		x /= 10
	}

	i := 0
	j := len(nums) - 1

	for i < j {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}

	return true
}
