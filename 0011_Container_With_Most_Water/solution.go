package main

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxVol := 0

	for i < j {
		var vol int
		if height[i] > height[j] {
			vol = (j - i) * height[j]
			j--
		} else {
			vol = (j - i) * height[i]
			i++
		}

		if vol > maxVol {
			maxVol = vol
		}
	}

	return maxVol
}
