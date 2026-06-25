package tasks

import (
	"fmt"
)

func FindTheDuplicateNumberRun() {
	fmt.Println("Task: 287. Find the Duplicate Number")
	fmt.Println("Link: https://leetcode.com/problems/find-the-duplicate-number/description/")
	fmt.Println("")

	var result int
	p := []int{1, 3, 4, 2, 2}
	fmt.Println("Input: ", p)
	result = findDuplicate(p)
	fmt.Println("Output: ", result)

	p = []int{3, 1, 3, 4, 2}
	fmt.Println("Input: ", p)
	result = findDuplicate(p)
	fmt.Println("Output: ", result)
}

func findDuplicate(nums []int) int {

    left, right := nums[0], nums[0]

    for {
        left = nums[left]
        right = nums[nums[right]]
        if left == right {
             break
        }
    }

    left = nums[0]
    for left != right {
        left = nums[left]
        right = nums[right]
    }

    return left
}