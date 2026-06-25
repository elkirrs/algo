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

	p = []int{2, 3, 5, 0}
	fmt.Println("Input: ", p)
	result = productExceptSelf(p)
	fmt.Println("Output: ", result)
}

func productExceptSelf(nums []int) []int {
    n := len(nums)
	out := make([]int, n)

	prefix := 1
	for i := 0; i < n; i++ {
		out[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		out[i] *= suffix
		suffix *= nums[i]
	}

	return out
}

// Для каждого элемента нужно получить:
// 1. произведение всех элементов слева;
// 2. произведение всех элементов справа.
