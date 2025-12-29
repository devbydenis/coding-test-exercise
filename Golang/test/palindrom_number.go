package test

import "fmt"

func PalindromNumber(number int) bool { // 1234
	temp := number // 1234
	reverseNumber := 0
	for temp > 0 {
		ekor := temp % 10 // 4

		reverseNumber = (reverseNumber * 10) + ekor

		fmt.Println(reverseNumber)
		temp = temp / 10
	}

	return number == reverseNumber
}
