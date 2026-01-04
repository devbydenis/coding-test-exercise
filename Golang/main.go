package main

import (
	// "fmt"
	"coding-exercise-golang/test"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Runtime: %v", time.Since(start))
	}()

	// ======================================================================
	// Two Sum
	// ======================================================================
	// var nums []int = []int{2, 7, 11, 15}
	// var target int = 9
	// fmt.Println(TwoSum(nums, target))

	// ======================================================================
	// Valid Parentheses
	// ======================================================================
	// var s string = "()[]{}"
	// fmt.Println(IsValid(s))

	// ======================================================================
	// Running Sum of 1d Array
	// ======================================================================

	// fmt.Println(RunningSum(nums))

	// ======================================================================
	// Palindrom Number
	// ======================================================================
	// fmt.Println(test.PalindromNumber(12321))

	// ======================================================================
	// Contain Duplicate
	// ======================================================================
	// fmt.Println(test.ContainDuplicate([]int{1, 2, 3, 4, 1}))
	// fmt.Println(test.ContainDuplicate([]int{1, 2, 3, 4}))

	// ======================================================================
	// Valid Palindrome
	// ======================================================================
	// fmt.Println(test.ValidPalindrome("bahlil anjing"))
	// fmt.Println(test.ValidPalindrome("Kasur rusak"))

	// ======================================================================
	// Longest Substring Without Repeating Characters
	// ======================================================================
	// fmt.Println(test.LengthOfLongestSubstring(""))
	// fmt.Println(test.LengthOfLongestSubstring(" "))
	// fmt.Println(test.LengthOfLongestSubstring("bbbbbbb"))
	// fmt.Println(test.LengthOfLongestSubstring("pwwkew"))
	// fmt.Println(test.LengthOfLongestSubstring("abcabcbb"))
	// fmt.Println(test.LengthOfLongestSubstring("bcbbaxn"))
	// fmt.Println(test.LengthOfLongestSubstring("pww"))

	// ======================================================================
	// Fibonacci Number
	// ======================================================================
	fmt.Println(test.Fib(2))
	fmt.Println(test.Fib(3))
	fmt.Println(test.Fib(4))
	// fmt.Println(test.Fib(10))
	// ======================================================================
}
