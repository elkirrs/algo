package tasks

import (
	"fmt"
	"sort"
)

func NumberOfDistinctAveragesRun() {
	fmt.Println("Task: 2465. Number of Distinct Averages")
	fmt.Println("Link: https://leetcode.com/problems/number-of-distinct-averages/description/")
	fmt.Println("")

	var result int
	p := []int{9,5,7,8,7,9,8,2,0,7}
	fmt.Println("Input: ", p)
	result = distinctAverages(p)
	fmt.Println("Output: ", result)

	p = []int{-1, 1, 0, -3, 3}
	fmt.Println("Input: ", p)
	result = distinctAverages(p)
	fmt.Println("Output: ", result)
}

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	hT := make(map[int]struct{})

	for l < r {
		sum := nums[l] + nums[r]
		hT[sum] = struct{}{}
		l++
		r--
	}

	return len(hT)
}