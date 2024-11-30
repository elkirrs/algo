package algo

import (
	"fmt"
	"sort"
)

func BinarySearch() {
	var list = []int{1, 4, 5, 8, 5, 1, 2, 7, 5, 2, 11}
	var target = 7
	sort.Ints(list)
	binarySearchRun(list, target)
}

// BinarySearch O(log2n)
func binarySearchRun(list []int, target int) int {
	var mid int = -1
	var start int = 0
	var end int = len(list) - 1
	var iterations int = 0

	for start <= end {
		iterations++
		mid = start + (end-start)/2
		if list[mid] == target {
			break
		} else if list[mid] < target {
			start = mid + 1 // search in left
		} else {
			end = mid - 1 // search in right
		}
	}
	fmt.Printf("binarySearch: value: %d index: %d iterations: %d \n", target, mid, iterations)

	return mid
}

// RecursiveBinarySearch
func RecursiveBinarySearch() int {
	var list = []int{1, 4, 5, 8, 5, 1, 2, 7, 5, 2, 11}
	var target = 7
	iterations := 0
	sort.Ints(list)
	result := recursiveBinarySearchRun(list, target, 0, len(list)-1, &iterations)
	fmt.Printf("recursiceBinarySearch: value: %d index: %d iterations: %d \n", target, result, iterations)
	return result
}

func recursiveBinarySearchRun(list []int, target, low, high int, iterations *int) int {
	*iterations++

	if low > high {
		return -1
	}

	mid := low + (high-low)/2

	if target == list[mid] {
		return mid
	}

	if target < list[mid] {
		return recursiveBinarySearchRun(list, target, low, mid-1, iterations)
	}

	return recursiveBinarySearchRun(list, target, mid+1, high, iterations)
}
