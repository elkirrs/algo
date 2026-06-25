package tasks

import (
	"fmt"
)

func TwoSumIIInputArrayIsSortedRun() {
	fmt.Println("167. Two Sum II - Input Array Is Sorted")
	fmt.Println("https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/")
	fmt.Println("")

	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("Input: numbers =", numbers, ", target =", target)
	result := twoSumIIInputArrayIsSorted(numbers, target)
	fmt.Println("Output:", result)
	fmt.Println("")
}

func twoSumIIInputArrayIsSorted(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1

    for left < right && numbers[left] + numbers[right] != target {        
        if numbers[left] + numbers[right] > target {
            right--
        } else {
            left++
        }
    }

    return []int{left+1, right+1}
}