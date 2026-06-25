package tasks

import (
	"fmt"
)

func FindMinimumInRotatedSortedArrayRun() {
	fmt.Println("153. Find Minimum in Rotated Sorted Array")
	fmt.Println("https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/")
	fmt.Println("")

	nums := []int{3, 4, 5, 1, 2}
	fmt.Println("Input: nums =", nums)
	result := findMin(nums)
	fmt.Println("Output:", result)
	fmt.Println("")
	
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("Input: nums =", nums)
	result = findMin(nums)
	fmt.Println("Output:", result)
	fmt.Println("")

	nums = []int{11, 13, 15, 17}
	fmt.Println("Input: nums =", nums)
	result = findMin(nums)
	fmt.Println("Output:", result)
	fmt.Println("")
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
