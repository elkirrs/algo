package tasks

import (
	"fmt"
)

func FindFirstAndLastPositionOfElementInSortedArrayRun() {
	fmt.Println("34. Find First and Last Position of Element in Sorted Array")
	fmt.Println("https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/")
	fmt.Println("")

	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println("Input:", nums, "Target:", target)
	result := searchRange(nums, target)
	fmt.Println("Output:", result)
}

func searchRange(nums []int, target int) []int {
    result := []int{-1, -1}
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            result[0] = mid
            right = mid - 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    left, right = 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            result[1] = mid
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return result
}