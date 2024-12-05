package binarySearch

import "fmt"

func SearchMatrix(matrix [][]int, target int) bool {
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

	//row, col := 0, len(matrix[0])-1
	//for row < len(matrix) && col >= 0 {
	//	iterations++
	//	if matrix[row][col] == target {
	//		fmt.Printf("iter: %d\n", iterations)
	//		return true
	//	} else if matrix[row][col] < target {
	//		row++
	//	} else {
	//		col--
	//	}
	//}

	return false
}

func SearchMatrixTwo(matrix [][]int, target int) bool {

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

func GenerateMatrix(sizeMatrix int) [][]int {
	matrix := make([][]int, sizeMatrix)

	for i := range matrix {
		matrix[i] = make([]int, sizeMatrix)
	}

	value := 1
	for i := 0; i < sizeMatrix; i++ {
		for j := 0; j < sizeMatrix; j++ {
			matrix[i][j] = value
			value++
		}
	}
	return matrix
}

func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
