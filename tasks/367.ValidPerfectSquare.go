package tasks

import (
	"fmt"
)

func ValidPerfectSquareRun() {
	fmt.Println("367. Valid Perfect Square")
	fmt.Println("https://leetcode.com/problems/valid-perfect-square/description/")
	fmt.Println("")

	num := 16

	fmt.Println("Input:")
	fmt.Println("num:", num)
	result := isPerfectSquare(num)
	fmt.Println("Result:", result)
}

func isPerfectSquare(num int) bool {
    left, right := 1, num

    for left <= right {
        mid := left + (right - left) / 2
        square := mid * mid
        if square == num {
            return true
        } else if square < num {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return false
}