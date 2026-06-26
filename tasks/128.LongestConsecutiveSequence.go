package tasks

import (
	"fmt"
	"slices"
)

func LongestConsecutiveSequenceRun() {
	fmt.Println("128. Longest Consecutive Sequence")
	fmt.Println("https://leetcode.com/problems/longest-consecutive-sequence/description/")
	fmt.Println("")

	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("Input:", nums)
	res := longestConsecutive(nums)
	fmt.Println("Output:", res)

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println("Input:", nums)
	res = longestConsecutive(nums)
	fmt.Println("Output:", res)
}

func longestConsecutive(nums []int) int {
	res := 0 
	if len(nums) == 0 {
    	return 0
    }

	counter := 1
	res = 1
	slices.Sort(nums)

	for i := 1; i < len(nums); i++ {

		if nums[i-1]+1 == nums[i] {
			counter++
		} else {
			if nums[i-1] != nums[i] {
				counter = 1
			}
		}

		if res < counter {
			res = counter
		}

	}

	return res
}
