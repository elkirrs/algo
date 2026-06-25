package tasks

import (
	"fmt"
)

func IntersectionOfTwoArraysIIRun() {
	fmt.Println("350. Intersection of Two Arrays II")
	fmt.Println("https://leetcode.com/problems/intersection-of-two-arrays-ii/description/")
	fmt.Println("")

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println("Input:")
	fmt.Println("nums1:", nums1)
	fmt.Println("nums2:", nums2)
	result := intersect(nums1, nums2)
	fmt.Println("Result:", result)
}


func intersect(nums1 []int, nums2 []int) []int {
    var mapArr, result = make(map[int]int), []int{}

    for _, num := range nums1 {
        mapArr[num]++
    }

    for _, num := range nums2 {
        if mapArr[num] > 0 {
            result = append(result, num)
            mapArr[num]--
        }
        
    }

    return result
}