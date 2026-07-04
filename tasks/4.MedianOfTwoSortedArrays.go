package tasks

import (
	"fmt"
	"math"
)

func MedianOfTwoSortedArraysRun() {
	fmt.Println("4. Median of Two Sorted Arrays")
	fmt.Println("https://leetcode.com/problems/median-of-two-sorted-arrays/description/")
	fmt.Println("")

	medianOfTwoSortedArraysPrint([]int{1, 3}, []int{2})
	medianOfTwoSortedArraysPrint([]int{1, 2,}, []int{3, 4})
	medianOfTwoSortedArraysPrint([]int{1,2,3,4,5}, []int{6,7,8,9,10,11,12,13,14,15,16,17})
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m

	for left <= right {
		cut1 := left + (right - left) /2
		cut2 := (m+n+1)/2 - cut1

		var l1, r1, l2, r2 int

		if cut1 == 0 {
			l1 = math.MinInt
		} else {
			l1 = nums1[cut1-1]
		}

		if cut1 == m {
			r1 = math.MaxInt
		} else {
			r1 = nums1[cut1]
		}

		if cut2 == 0 {
			l2 = math.MinInt
		} else {
			l2 = nums2[cut2-1]
		}

		if cut2 == n {
			r2 = math.MaxInt
		} else {
			r2 = nums2[cut2]
		}

		if l1 <= r2 && l2 <= r1 {
			if (n+m)%2 == 1 {
				return float64(max(l1, l2))
			}

		return float64(max(l1, l2) + min(r1, r2)) / 2.0
		}

		if l1 > r2 {
			right = cut1 - 1
		} else {
			left = cut1 +1
		}

	}

	return 0
}

func medianOfTwoSortedArraysPrint(nums1 []int, nums2 []int) {
	fmt.Println("Input:", "nums1 =", nums1, "nums2 =", nums2)
	out := findMedianSortedArrays(nums1, nums2)
	fmt.Println("Output:", out)
}
