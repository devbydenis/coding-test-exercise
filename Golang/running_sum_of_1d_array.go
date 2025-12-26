package main

func RunningSum(nums []int) []int { // [1,2,3,4]
	var result []int

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result = append(result, nums[i])
			continue
		}

		var temp int = 0
		for _, n := range nums[:i + 1] {
			temp += n
		}
		result = append(result, temp)
	}

	return result
}