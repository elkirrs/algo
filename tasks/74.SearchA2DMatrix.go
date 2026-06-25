package tasks

import (
	"fmt"
) 

func SearchA2DMatrixRun() {
	fmt.Println("74. Search a 2D Matrix")
	fmt.Println("https://leetcode.com/problems/search-a-2d-matrix/")
	fmt.Println("")

	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3
	fmt.Println("Input:")
	fmt.Println("[1,  3,  5,   7]")
	fmt.Println("[10, 11, 16, 20]")
	fmt.Println("[23, 30, 34, 60]")
	fmt.Println("Target:", target)
	result := searchMatrix(matrix, target)
	fmt.Println("Output:", result)
}

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }

	rowDown, rowUp := 0, len(matrix)-1
	colLeft, colRight := 0, len(matrix[0])-1

	for rowDown <= rowUp {
		midRow := rowDown + (rowUp-rowDown)/2
		if matrix[midRow][colLeft] <= target && matrix[midRow][colRight] >= target {
			for colLeft <= colRight {
				midCol := colLeft + (colRight-colLeft)/2
				if matrix[midRow][midCol] == target {
					return true
				} else if matrix[midRow][midCol] < target {
					colLeft = midCol + 1
				} else {
					colRight = midCol - 1
				}
			}
		} else if matrix[midRow][colLeft] > target {
			rowUp = midRow - 1
		} else {
			rowDown = midRow + 1
		}
	}

    return false
}