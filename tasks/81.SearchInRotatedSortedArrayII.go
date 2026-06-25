package tasks

import (
	"fmt"
)

func SearchInRotatedSortedArrayIIRun() {
	fmt.Println("81. Search in Rotated Sorted Array II")
	fmt.Println("https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/")
	fmt.Println("")

	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	fmt.Println("Input: nums =", nums, ", target =", target)
	result := searchII(nums, target)
	fmt.Println("Output:", result)
	fmt.Println("")

	nums = []int{2, 5, 6, 0, 0, 1, 2}
	target = 3
	fmt.Println("Input: nums =", nums, ", target =", target)
	result = searchII(nums, target)
	fmt.Println("Output:", result)
	fmt.Println("")
	nums = []int{1, 0, 1, 1, 1}
	target = 0
	fmt.Println("Input: nums =", nums, ", target =", target)
	result = searchII(nums, target)
	fmt.Println("Output:", result)
	fmt.Println("")


}

func searchII(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		if nums[left] == nums[mid] && nums[mid] == nums[right] {
            left++
            right--
        } else if nums[left] <= nums[mid] {
            if nums[left] <= target && nums[mid] > target {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
        }

	}

	return false
}
