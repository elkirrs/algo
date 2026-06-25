package tasks

import (
	"fmt"
)

func SearchA2DMatrixIIRun() {
	fmt.Println("240. Search a 2D Matrix II")
	fmt.Println("https://leetcode.com/problems/search-a-2d-matrix-ii/description/")
	fmt.Println("")

	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
	}
	target := 5
	fmt.Println("Input: matrix =")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("target =", target)
	fmt.Println("Output:", searchMatrix2(matrix, target))
}

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := len(matrix)-1, 0

	for row >= 0 && col <= len(matrix[0])-1 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			col++
		} else {
			row--
		}
	}

	return false
}
