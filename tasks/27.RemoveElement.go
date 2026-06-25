package tasks

import (
	"fmt"
)

func RemoveElementRun() {
	fmt.Println("27. Remove Element")
	fmt.Println("https://leetcode.com/problems/remove-element/")
	fmt.Println("")

	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println("Input:", nums, "Value to remove:", val)
	result := removeElement(nums, val)
	fmt.Println("Output:", result)
}

func removeElement(nums []int, val int) int {
    s := 0
    for f := 0; f < len(nums); f++ {
        if nums[f] != val {
            nums[s] = nums[f]
            s++
        }
    }
    return s
}
