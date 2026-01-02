package test

func ContainDuplicate(nums []int) bool { // [1, 2, 3, 4]
	m := make(map[int]int)

	for _, num := range nums {
		if _, found := m[num]; found {
			return true
		}

		m[num] = num
	}

	return false
}
