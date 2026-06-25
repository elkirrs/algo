package tasks

import "fmt"

func ProductofArrayExceptSelfRun() {
	fmt.Println("Task: 238. Product of Array Except Self")
	fmt.Println("Link: https://leetcode.com/problems/product-of-array-except-self/description/")
	fmt.Println("")

	var result []int
	p := []int{1, 2, 3, 4}
	fmt.Println("Input: ", p)
	result = productExceptSelf(p)
	fmt.Println("Output: ", result)

	p = []int{-1, 1, 0, -3, 3}
	fmt.Println("Input: ", p)
	result = productExceptSelf(p)
	fmt.Println("Output: ", result)
}

func productExceptSelf(nums []int) []int {
	var result = make([]int, len(nums))


	return result
}