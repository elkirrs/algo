package tasks 

import (
	"fmt"
)

func IntersectionOfTwoArraysRun() {
	fmt.Println("349. Intersection of Two Arrays")
	fmt.Println("https://leetcode.com/problems/intersection-of-two-arrays/")
	fmt.Println("")

	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Println("Input: nums1 =", nums1, ", nums2 =", nums2)
	result := intersection(nums1, nums2)
	fmt.Println("Output: ", result)
}

func intersection(nums1 []int, nums2 []int) []int {
    result := []int{}
    mapArr := make(map[int]int)

    for _, num := range nums1 {
        mapArr[num] = 1
    }

    for _, num := range nums2 {
        if mapArr[num] > 0 {
            result = append(result, num)
            mapArr[num]--
        }
    }

    return result

}