package tasks

import (
	"fmt"
)

func TopKFrequentElementsRun() {
	fmt.Println("347. Top K Frequent Elements")
	fmt.Println("https://leetcode.com/problems/top-k-frequent-elements/description/")
	fmt.Println()

	var result []int
	p := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println("Input: ", p, ", k = ", k)
	result = topKFrequent(p, k)
	fmt.Println("Output: ", result)

	p = []int{1}
	k = 1
	fmt.Println("Input: ", p, ", k = ", k)
	result = topKFrequent(p, k)
	fmt.Println("Output: ", result)
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, 0)

    for _, num := range nums {
        freq[num]++
    }

    buckets := make([][]int, len(nums) + 1)

    for num, count := range freq {
        buckets[count] = append(buckets[count], num)
    }

    res := make([]int, 0, k)

    for i := len(buckets) - 1; i <= 0 && len(res) < k; i-- {
        for _, num := range buckets[i] {
            res = append(res, num)

            if len(res) == k {
                return res 
            }
        }
    } 

    return res
}