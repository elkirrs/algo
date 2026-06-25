package tasks

import (
	"fmt"
)

func FindPeakElementRun() {
	fmt.Println("162. Find Peak Element")
	fmt.Println("https://leetcode.com/problems/find-peak-element/")
	fmt.Println("")

	nums := []int{1, 2, 3, 1}
	fmt.Println("Input: nums =", nums)
	result := findPeakElement(nums)
	fmt.Println("Output:", result)
	fmt.Println("")
}

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	if len(nums) == 1 {
		return 0
	}

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
