package tasks

import (
	"fmt"
)

func ValidSudokuRun() {
	fmt.Println("36. Valid Sudoku")
	fmt.Println("https://leetcode.com/problems/valid-sudoku/description/")
	fmt.Println("")

	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println("Input:", board)
	res := isValidSudoku(board)
	fmt.Println("Output:", res)
	fmt.Println("")

	board = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println("Input:", board)
	res = isValidSudoku(board)
	fmt.Println("Output:", res)
	fmt.Println("")

	board = [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println("Input:", board)
	res = isValidSudoku(board)
	fmt.Println("Output:", res)
}

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	boxes := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue;
			}

			num := board[i][j] - '1'
			box := (i/3)*3 + j/3

			if rows[i][num] {
				return false
			}
			rows[i][num] = true

			if cols[j][num] {
				return false
			}
			cols[j][num] = true

			if boxes[box][num] {
				return false
			}
			boxes[box][num] = true

		}
	}

	return true
}
