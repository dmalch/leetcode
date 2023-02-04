package main

func twoSum(nums []int, target int) []int {
	knownIndexes := make(map[int]int, len(nums))

	for i, num := range nums {
		if val, ok := knownIndexes[num]; ok {
			return []int{i, val}
		}

		knownIndexes[target-num] = i
	}

	return nil
}
