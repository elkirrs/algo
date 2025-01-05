package twoPointers

//nums = []int{-10, -8, -2, 1, 2, 5, 6}
//nums = []int{3, 24, 50, 79, 88, 150, 345}
//target = 0

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	var result []int

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return result
}
