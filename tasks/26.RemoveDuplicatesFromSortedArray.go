package tasks

import (
	"fmt"
)

func RemoveDuplicatesFromSortedArrayRun() {
	fmt.Println("26. Remove Duplicates from Sorted Array")
	fmt.Println("https://leetcode.com/problems/remove-duplicates-from-sorted-array/")
	fmt.Println("")

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Input:", nums)
	result := removeDuplicates(nums)
	fmt.Println("Output:", result)
}

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    l := 0
    for r := 1; r < len(nums); r++ {
        if nums[r] != nums[l] {
            l++
            nums[l] = nums[r]
        }
    }
    return l+1

}