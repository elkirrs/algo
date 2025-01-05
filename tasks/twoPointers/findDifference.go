package twoPointers

//nums1 = [1,2,3]
//nums2 = [2,4,6]

func findDifference(nums1 []int, nums2 []int) [][]int {
	var row1, row2 []int
	mapArr := make(map[int]bool)
	added := make(map[int]bool)

	for _, num := range nums1 {
		mapArr[num] = true
	}

	for _, num := range nums2 {
		if _, ok := mapArr[num]; !ok && !added[num] {
			row2 = append(row2, num)
			added[num] = true
		} else {
			mapArr[num] = false
		}

	}

	for num, val := range mapArr {
		if val == true {
			row1 = append(row1, num)
		}
	}

	return [][]int{row1, row2}
}

//func findDifference(nums1 []int, nums2 []int) [][]int {
//	mapNums1 := make(map[int]struct{})
//	mapNums2 := make(map[int]struct{})
//
//	output := [][]int{{}, {}}
//
//	for _, value := range nums1 {
//		mapNums1[value] = struct{}{}
//	}
//
//	for _, value := range nums2 {
//		mapNums2[value] = struct{}{}
//	}
//
//	for key := range mapNums1 {
//		if _, ok := mapNums2[key]; !ok {
//			output[0] = append(output[0], key)
//		}
//	}
//
//	for key := range mapNums2 {
//		if _, ok := mapNums1[key]; !ok {
//			output[1] = append(output[1], key)
//		}
//	}
//
//	return output
//}
