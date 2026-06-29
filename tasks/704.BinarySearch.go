package tasks

import (
	"fmt"
)

func BinarySearchRun() {
	fmt.Println("704. Binary Search")
	fmt.Println("https://leetcode.com/problems/binary-search/description/")
	fmt.Println("")

	binarySearchPrint([]int{-1, 0, 3, 5, 9, 12}, 9)
	binarySearchPrint([]int{-1, 0, 3, 5, 9, 12}, 2)
	binarySearchPrint([]int{-1, 0, 3, 5, 9, 12}, 13)
}

func searchBinary(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func binarySearchPrint(nums []int, target int) {
	fmt.Printf("Input: nums = %v, target = %v\n", nums, target)
	result := searchBinary(nums, target)
	fmt.Printf("Output: %v\n", result)
	fmt.Println("")
}
