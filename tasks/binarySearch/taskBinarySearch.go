package binarySearch

import "fmt"

func TasksRun() {
	var target int
	var nums []int
	var result int
	var resultBool bool
	var resultArr []int
	var matrix [][]int

	//firstBadVersion
	target = 5
	result = FirstBadVersion(target)
	fmt.Printf("firstBadVersion: target: %d result: %d\n", target, result)

	//isPerfectSquare
	target = 25
	resultBool = IsPerfectSquare(target)
	fmt.Printf("isPerfectSquare: target: %d is: %t\n", target, resultBool)

	//searchInsertPosition
	nums = []int{1, 3, 5, 6}
	target = 5
	result = SearchInsertPosition(nums, target)
	fmt.Printf("searchInsertPosition: %v target: %d position: %d\n", nums, target, result)

	//mySqrt
	target = 8
	result = MySqrt(target)
	fmt.Printf("mySqrt: target: %d result: %d\n", target, result)

	//searchInRotatedSortedArray
	target = 0
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	result = SearchInRotatedSortedArray(nums, target)
	fmt.Printf("searchInRotatedSortedArray: %v target: %d position: %d\n", nums, target, result)

	//peakIndexInMountainArray
	nums = []int{4, 5, 6, 7, 3, 2, 1, 0}
	result = PeakIndexInMountainArray(nums)
	fmt.Printf("peakIndexInMountainArray: %v position: %d\n", nums, result)

	//findFirstLastPositionElement
	nums = []int{5, 7, 7, 8, 8, 10}
	target = 8
	resultArr = SearchRange(nums, target)
	fmt.Printf("findFirstLastPositionElement: %v target: %d position: %v\n", nums, target, resultArr)

	//searchMatrix
	matrix = GenerateMatrix(10)
	target = 38
	resultBool = SearchMatrix(matrix, target)
	PrintMatrix(matrix)
	fmt.Printf("searchMatrix: target: %d position: %t\n", target, resultBool)

	//searchMatrix2
	matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target = 14
	resultBool = SearchMatrixTwo(matrix, target)
	PrintMatrix(matrix)
	fmt.Printf("searchMatrix2: target: %d position: %t\n", target, resultBool)

}
