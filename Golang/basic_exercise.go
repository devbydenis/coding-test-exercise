package main

import (
	"fmt"
)

func DeleteElement() {
	nums := []int{1, 2, 3, 4, 5}
	deleteIndex := 2

	nums = append(nums[:deleteIndex], nums[deleteIndex+1:]...)
	fmt.Println(nums)
}

func CreateSlice() {
	var exampleSlice = make([]int, 7, 10)

	for i := 0; i < 7; i++ {
		exampleSlice[i] = i + 1
	}

	fmt.Println(exampleSlice)
}

func CopySlice() {
	slice1 := []int{1, 2, 3, 4, 5}
	newSlice := make([]int, len(slice1))
	copy(newSlice, slice1)

	fmt.Println(newSlice)
}

func ReverseSlice() {
	nums := []int{1, 2, 3, 4, 5}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	fmt.Println(nums)
}

func MergeTwoSlice() {
	a := []int{1, 2, 3}
	b := []int{4, 5}
	merged := make([]int, 0, len(a)+len(b))

	merged = append(merged, a...)
	merged = append(merged, b...)

	fmt.Println(merged)
}

func FrequentCount() {
	nums := []int{2, 2, 3, 3, 3, 4, 5, 5}
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	for i, v := range freq {
		fmt.Printf("%d:%d \n", i, v)
	}
}

func MostFrequent() {
	nums := []int{2, 2, 3, 3, 3, 4, 5, 5}
	freq := make(map[int]int)
	maxValue, maxFreq := nums[0], 0

	for _, num := range nums {
		freq[num]++
		if freq[num] > maxFreq {
			maxFreq = freq[num]
			maxValue = num
		}
	}

	fmt.Println(maxValue)
}