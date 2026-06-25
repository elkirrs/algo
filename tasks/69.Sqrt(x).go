package tasks

import (
	"fmt"
)

func SQRTxRun() {
	fmt.Println("69. Sqrt(x)")
	fmt.Println("https://leetcode.com/problems/sqrtx/")
	fmt.Println("")

	x := 8
	fmt.Println("Input:", x)
	result := mySqrt(x)
	fmt.Println("Output:", result)
}

func mySqrt(x int) int {
    left, right := 0, x

    for left <= right {
        mid := left + (right-left)/2
        square := mid*mid
        if square == x {
            return mid
        } else if square < x {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return right
}