package tasks

import (
	"fmt"
)

func SearchInRotatedSortedArrayRun() {
	fmt.Println("33. Search in Rotated Sorted Array")
	fmt.Println("https://leetcode.com/problems/search-in-rotated-sorted-array/")
	fmt.Println("")

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println("Input:", nums, "Target:", target)
	result := search(nums, target)
	fmt.Println("Output:", result)
}

func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }

        if nums[left] <= nums[mid] {
            if target >= nums[left] && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            if target > nums[mid] && target <= nums[right]{
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }

    return -1
}