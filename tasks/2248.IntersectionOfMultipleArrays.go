package tasks

import (
	"fmt"
)

func IntersectionOfMultipleArraysRun() {
	fmt.Println("2248. Intersection of Multiple Arrays")
	fmt.Println("https://leetcode.com/problems/intersection-of-multiple-arrays/description/")
	fmt.Println("")

	nums := [][]int{
		{3, 1, 2, 4, 5},
		{1, 2, 3, 4},
		{3, 4, 5, 6},
	}

	fmt.Println("Input:")
	fmt.Println("nums:", nums)
	result := intersectionM(nums)
	fmt.Println("Result:", result)
}

func intersectionM(nums [][]int) []int {
	var result []int
	mapArr := make([]int, 1001)
	lenNums := len(nums)

	for _, arr := range nums {
		for _, num := range arr {
			mapArr[num]++;
		}
	}

    for num, counter := range mapArr {
        if counter == lenNums {
			result = append(result, num)
		}
    }

	return result
}