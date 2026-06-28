package tasks

import (
	"fmt"
)

func ContainerWithMostWaterRun() {
	fmt.Println("11. Container With Most Water")
	fmt.Println("https://leetcode.com/problems/container-with-most-water/description/")
	fmt.Println("")

	containerWithMostWaterPrint([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	containerWithMostWaterPrint([]int{1, 1})
	containerWithMostWaterPrint([]int{8, 7, 2, 1})
	containerWithMostWaterPrint([]int{1, 2, 3, 1000, 9})
}

func containerWithMostWater(height []int) int {
	var maxArea int

	left := 0
	right := len(height) - 1

	for left < right {
		min := min(height[left], height[right])
		area := min * (right - left)

		if maxArea < area {
			maxArea = area
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}

func containerWithMostWaterPrint(in []int) {
	fmt.Println("Input:", in)
	out := containerWithMostWater(in)
	fmt.Println("Output:", out)
	fmt.Println("")
}
