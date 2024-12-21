package binarySearch

//nums = []int{-10, -8, -2, 1, 2, 5, 6}
//nums = []int{3, 24, 50, 79, 88, 150, 345}
//target = 0

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	result := []int{-1, -1}

	for left <= right {
		mid := left + (right-left)/2

		if numbers[left]+numbers[right] == target {
			result[0] = left + 1
			result[1] = right + 1
			return result
		} else if numbers[mid] > target {
			if numbers[left]+numbers[mid] < target {
				left++
			} else {
				right = mid
			}
		} else if numbers[right]+numbers[left] < target {
			left++
		} else {
			right--
		}
	}

	return result
}
