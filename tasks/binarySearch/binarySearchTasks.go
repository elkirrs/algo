package binarySearch

import fmt "fmt"

func TusksRun() {
	//firstBadVersion
	fbv := firstBadVersion(5)
	fmt.Printf("firstBadVersion: %d\n", fbv)

	//isPerfectSquare
	num := 25
	vps := isPerfectSquare(num)
	fmt.Printf("isPerfectSquare: %d is: %t\n", num, vps)

	//searchInsertPosition
	sipNums := []int{1, 3, 5, 6}
	target := 5
	sip := searchInsertPosition(sipNums, target)
	fmt.Printf("searchInsertPosition: %v position: %d\n", sipNums, sip)

}
