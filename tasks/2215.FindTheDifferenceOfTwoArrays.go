package tasks

import "fmt"

func FindTheDifferenceOfTwoArraysRun() {
	fmt.Println("Task: 2215. Find the Difference of Two Arrays")
	fmt.Println("Link: https://leetcode.com/problems/find-the-difference-of-two-arrays/description/")
	fmt.Println("")

	var result [][]int
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	fmt.Println("Input: ", nums1, nums2)
	result = findDifference(nums1, nums2)
	fmt.Println("Output: ", result)
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	var row1, row2 []int

	mapArr := make(map[int]bool)
    added := make(map[int]bool)

	for _, num := range nums1 {
		mapArr[num] = true
	}

	for _, num := range nums2 {
		if _, ok := mapArr[num]; !ok && !added[num] {
			row2 = append(row2, num)
            added[num] = true
		} else {
            mapArr[num] = false
        }
		
	}

	for num, val := range mapArr {
		if val == true {
			row1 = append(row1, num)
		}
	}

	return [][]int{row1, row2}
}