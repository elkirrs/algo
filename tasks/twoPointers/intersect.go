package twoPointers

// nums1 = [4,9,5]
//nums2 = [9,4,9,8,4]

func intersect(nums1 []int, nums2 []int) []int {
	var result []int
	var mapArr = make(map[int]int)

	for _, num := range nums1 {
		mapArr[num]++
	}

	for _, num := range nums2 {
		if mapArr[num] > 0 {
			result = append(result, num)
		}
		mapArr[num]--
	}

	return result
}
