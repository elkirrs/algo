package twoPointers

//nums = [9,5,7,8,7,9,8,2,0,7]

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	hT := make(map[int]struct{})

	for l < r {
		sum := nums[l] + nums[r]
		hT[sum] = struct{}{}
		l++
		r--
	}

	return len(hT)
}
