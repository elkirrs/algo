package binarySearch

import fmt "fmt"

func TasksRun() {
	//firstBadVersion
	fbvTarget := 5
	fbv := firstBadVersion(fbvTarget)
	fmt.Printf("firstBadVersion: target: %d result: %d\n", fbvTarget, fbv)

	//isPerfectSquare
	vpsTarget := 25
	vps := isPerfectSquare(vpsTarget)
	fmt.Printf("isPerfectSquare: target: %d is: %t\n", vpsTarget, vps)

	//searchInsertPosition
	sipNums := []int{1, 3, 5, 6}
	sipTarget := 5
	sip := searchInsertPosition(sipNums, sipTarget)
	fmt.Printf("searchInsertPosition: %v target: %d position: %d\n", sipNums, sipTarget, sip)

	//mySqrt
	msTarget := 8
	ms := mySqrt(msTarget)
	fmt.Printf("mySqrt: target: %d result: %d\n", msTarget, ms)
}
