package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nums []int
	nums = append(nums, 3)
	nums = append(nums, 3)
	fmt.Println(twoSum(nums, 6))

	fmt.Println(isPalindrome(11))

	kata := 10
	angka := strconv.Itoa(kata)
	kiri := 0
	kanan := len(angka) - 1

	fmt.Println(angka[kiri])
	fmt.Println(angka[kanan])
}

func isPalindrome(x int) bool {
	newSTRING := strconv.Itoa(x)
	left := 0
	right := len(newSTRING) - 1
	for left < right {
		if newSTRING[left] != newSTRING[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func twoSum(nums []int, target int) []int {
	var twoSums []int
	for i := 0; i <= len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				twoSums = append(twoSums, nums[i])
				return []int{i, j}
			}
		}
	}
	return nil
}

//func palindrome(nums int) bool {
//	for i := nums - 1; i > 0; i-- {
//
//	}
//}
