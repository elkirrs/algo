package twoPointers

//nums1 := []int{1, 2, 2, 1}
//nums2 := []int{2, 2}

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
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
