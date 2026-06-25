package tasks

import (
	"fmt"
)

func SquaresOfASortedArrayRun() {
	fmt.Println("977. Squares of a Sorted Array")
	fmt.Println("https://leetcode.com/problems/squares-of-a-sorted-array/description/")
	fmt.Println("")

	arr := []int{-4, -1, 0, 3, 10}

	fmt.Println("Input:")
	fmt.Println("arr:", arr)
	result := sortedSquares(arr)
	fmt.Println("Result:", result)
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	slice := make([]int, len(nums))
	l, r := 0, n-1
	pos := n - 1

	for l <= r {
		lSquare := nums[l] * nums[l]
		rSquare := nums[r] * nums[r]

		if lSquare > rSquare {
			slice[pos] = lSquare
			l++
		} else {
			slice[pos] = rSquare
			r--
		}

		pos--
	}

	return slice
}
