package tasks

import (
	"fmt"
)

func FindMinimumInRotatedSortedArrayIIRun() {
	fmt.Println("154. Find Minimum in Rotated Sorted Array II")
	fmt.Println("https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/description/")
	fmt.Println("")

	nums := []int{1, 3, 5}
	fmt.Println("Input: nums =", nums)
	result := findMinII(nums)
	fmt.Println("Output:", result)
	fmt.Println("")
	
	nums = []int{2, 2, 2, 0, 1}
	fmt.Println("Input: nums =", nums)
	result = findMinII(nums)
	fmt.Println("Output:", result)
	fmt.Println("")
}

func findMinII(nums []int) int {
    left, right := 0, len(nums)-1

    for left < right {
        mid := left + (right - left) / 2

        if nums[left] == nums[mid] && nums[mid] == nums[right] {
            left++
            right--
        } else if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return nums[left]
}