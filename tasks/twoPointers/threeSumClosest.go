package twoPointers

import "sort"

//nums := []int{4, 0, 5, -5, 3, 3, 0, -4, -5} target := -2
//nums := []int{-1,2,1,-4} target := 1

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return sum
			}

			if abs(target-result) > abs(target-sum) {
				result = sum
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
