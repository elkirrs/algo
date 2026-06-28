package tasks

import (
	"fmt"
	"sort"
)

func ThreeSumRun() {
	fmt.Println("15. 3Sum")
	fmt.Println("https://leetcode.com/problems/3sum/")
	fmt.Println("")

	// threeSumPrint([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4})
    threeSumPrint([]int{0, 0, 0, 0})
}	


func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        left, right := i+1, len(nums)-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                left++
                right--
            } else if sum > 0 {
                right--
            } else {
                left++
            }
        }
    }

	return result
}

func threeSumPrint(in []int) {
	fmt.Println("Input:", in)
	output := threeSum(in)
	fmt.Println("Output:", output)
	fmt.Println("")
}