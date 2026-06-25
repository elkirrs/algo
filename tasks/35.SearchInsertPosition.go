package tasks

import (
	"fmt"
)

func SearchInsertPositionRun() {
	fmt.Println("35. Search Insert Position")
	fmt.Println("https://leetcode.com/problems/search-insert-position/")
	fmt.Println("")

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println("Input:", nums, "Target:", target)
	result := searchInsert(nums, target)
	fmt.Println("Output:", result)
}

func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return left
}