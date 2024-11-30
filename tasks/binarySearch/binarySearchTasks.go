package binarySearch

import fmt "fmt"

func TusksRun() {
	fbv := firstBadVersion(5)
	fmt.Printf("firstBadVersion: %d\n", fbv)

	num := 25
	vps := isPerfectSquare(num)
	fmt.Printf("isPerfectSquare: %d is: %t\n", num, vps)
}
