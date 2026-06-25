package tasks

import (
	"fmt"
)

func ContainsDuplicateRun() {
	fmt.Println("217. Contains Duplicate")
	fmt.Println("https://leetcode.com/problems/contains-duplicate/description/")
	fmt.Println()

	var result bool
	p := []int{1, 2, 3, 4, 5}
	fmt.Println("Input: ", p)
	result = containsDuplicate(p)
	fmt.Println("Output: ", result)

	p = []int{1, 2, 3, 4, 5, 1}
	fmt.Println("Input: ", p)
	result = containsDuplicate(p)
	fmt.Println("Output: ", result)
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}

	return false
}
