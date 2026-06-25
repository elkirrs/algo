package tasks

import (
	"fmt"
)

func FindPeakElementIIRun() {
	fmt.Println("1901. Find a Peak Element II")
	fmt.Println("https://leetcode.com/problems/find-a-peak-element-ii/description/")
	fmt.Println("")

	matrix := [][]int{
		{1, 4},
		{3, 2},
	}

	fmt.Println("Input:")
	fmt.Println("matrix:", matrix)
	result := findPeakGrid(matrix)
	fmt.Println("Result:", result)
}

func findPeakGrid(mat [][]int) []int {
	result := []int{-1, -1}

	if len(mat) == 0 && len(mat[0]) == 0 {
		return result
	}

	left, right := 0, len(mat[0])-1

	for left <= right {
		mid := left + (right-left)/2

		max := 0
		for i := 0; i < len(mat); i++ {
			if mat[i][mid] > mat[max][mid] {
				max = i
			}
		}

		if (mid == 0 || mat[max][mid] > mat[max][mid-1]) && (mid == len(mat[0])-1 || mat[max][mid] > mat[max][mid+1]) {
			result[0] = max
			result[1] = mid
			return result
		} else if mid > 0 && mat[max][mid] < mat[max][mid-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}
