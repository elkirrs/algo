package binarySearch

import "fmt"

func TasksRun() {
	//firstBadVersion
	fbvTarget := 5
	firstBadVersionResult := firstBadVersion(fbvTarget)
	fmt.Printf("firstBadVersion: target: %d result: %d\n", fbvTarget, firstBadVersionResult)

	//isPerfectSquare
	vpsTarget := 25
	isPerfectSquareResult := isPerfectSquare(vpsTarget)
	fmt.Printf("isPerfectSquare: target: %d is: %t\n", vpsTarget, isPerfectSquareResult)

	//searchInsertPosition
	sipNums := []int{1, 3, 5, 6}
	sipTarget := 5
	searchInsertPositionResult := searchInsertPosition(sipNums, sipTarget)
	fmt.Printf("searchInsertPosition: %v target: %d position: %d\n", sipNums, sipTarget, searchInsertPositionResult)

	//mySqrt
	msTarget := 8
	mySqrtResult := mySqrt(msTarget)
	fmt.Printf("mySqrt: target: %d result: %d\n", msTarget, mySqrtResult)

	//searchInRotatedSortedArray
	rsaTarget := 0
	rsaNums := []int{4, 5, 6, 7, 0, 1, 2}
	searchInRotatedSortedArrayResult := searchInRotatedSortedArray(rsaNums, rsaTarget)
	fmt.Printf("searchInRotatedSortedArray: %v target: %d position: %d\n", rsaNums, rsaTarget, searchInRotatedSortedArrayResult)

	//peakIndexInMountainArray
	pimaNums := []int{4, 5, 6, 7, 3, 2, 1, 0}
	peakIndexInMountainArrayResult := peakIndexInMountainArray(pimaNums)
	fmt.Printf("peakIndexInMountainArray: %v position: %d\n", pimaNums, peakIndexInMountainArrayResult)

}
