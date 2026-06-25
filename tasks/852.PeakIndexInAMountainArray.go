package tasks

import (
	"fmt"
)

func PeakIndexInAMountainArrayRun() {
	fmt.Println("852. Peak Index in a Mountain Array")
	fmt.Println("https://leetcode.com/problems/peak-index-in-a-mountain-array/description/")
	fmt.Println("")

	arr := []int{0, 1, 0}

	fmt.Println("Input:")
	fmt.Println("arr:", arr)
	result := peakIndexInMountainArray(arr)
	fmt.Println("Result:", result)
}

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
