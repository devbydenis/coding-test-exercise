package test

// a b c
// b b b b b
// p w w k e w
/*
p = 0
*/

func LengthOfLongestSubstring(s string) int { //  p w w k e w
	m := make(map[byte]int) // p: 0, w: 1, k: 3, e: 4
	left := 0               // _, _, 2,
	maxLength := 0          // 1, 2, _, 2, 3,

	for right := 0; right < len(s); right++ {
		char := s[right]

		if lastIndex, found := m[char]; found {  // cek apakah character yang sedang di looping itu duplicate? -> shrink
			if lastIndex+1 > left {
				left = lastIndex + 1
			}
		}

		m[char] = right  // set ke Map karna unique -> expand

		currentLength := right - left + 1  // panjang karakter yang ada didalam window saat ini

		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

// unique? -> expand
// duplicate? -> shrink

func findMaxSum(nums []int, k int) int { // 2, 1, 5, 1, 3, 2
	if len(nums) < k {
		return 0
	}

	currentWindowSum := 0
	maxSum := 0 // 8

	// 1. Itung jumlah jendela pertama (0 sampe k-1)
	for i := 0; i < k; i++ {
		currentWindowSum += nums[i]
	}
	maxSum = currentWindowSum

	// 2. Mulai "Sliding" atau geser jendelanya
	for i := k; i < len(nums); i++ {
		// Tambahin elemen baru (kanan), kurangin elemen lama (kiri)
		currentWindowSum += nums[i] - nums[i-k]

		// Update maxSum kalau nemu yang lebih gede
		if currentWindowSum > maxSum {
			maxSum = currentWindowSum
		}
	}

	return maxSum
}
