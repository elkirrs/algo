package tasks

import (
	"fmt"
)

func MergeSortedArrayRun() {
	fmt.Println("88. Merge Sorted Array")
	fmt.Println("https://leetcode.com/problems/merge-sorted-array/description/")
	fmt.Println("")

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3

	fmt.Println("Input: nums1 =", nums1, ", m =", m, ", nums2 =", nums2, ", n =", n)

	merge(nums1, m, nums2, n)

	fmt.Println("Output: nums1 =", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
