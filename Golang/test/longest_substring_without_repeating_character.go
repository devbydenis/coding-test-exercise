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
