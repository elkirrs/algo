package tasks

import (
	"fmt"
)

func TrappingRainWaterRun() {
	fmt.Println("42. Trapping Rain Water")
	fmt.Println("https://leetcode.com/problems/trapping-rain-water/description/")
	fmt.Println("")

	trappingRainWaterPrint([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}) // output 6
	trappingRainWaterPrint([]int{4, 2, 0, 3, 2, 5})                   // output 9
}

func trap(height []int) int {
	var maxArea int
	left, right := 0, len(height)-1

	for left < right {
		if height[left] < height[right] {
			minHeight := height[left]
			left++
			for left < right && height[left] < minHeight {
				maxArea += minHeight - height[left]
				left++
			}
		} else {
			minHeight := height[right]
			right--
			for left < right && height[right] < minHeight {
				maxArea += minHeight - height[right]
				right--
			}
		}
	}

	return maxArea
}

func trappingRainWaterPrint(in []int) {
	fmt.Println("Input:", in)
	out := trap(in)
	fmt.Println("Output:", out)
	fmt.Println("")
}
