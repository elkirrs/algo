package tasks

import (
	"fmt"
)

func RunningSumOf1dArrayRun() {
	fmt.Println("1480. Running Sum of 1d Array")
	fmt.Println("https://leetcode.com/problems/running-sum-of-1d-array/description/")
	fmt.Println("")

	arr := []int{1, 2, 3, 4}

	fmt.Println("Input:")
	fmt.Println("arr:", arr)
	result := runningSum(arr)
	fmt.Println("Result:", result)
}

func runningSum(nums []int) []int {
	for idx, num := range nums {
		if idx == 0 {
			nums[idx] = num
		} else {
			nums[idx] = nums[idx-1] + num
		}
	}

	return nums
}
