package twoPointers

//nums := [][]int{{7, 34, 45, 10, 12, 27, 13}, {27, 21, 45, 10, 12, 13}}

func intersectionMultiple(nums [][]int) []int {
	var result []int
	mapArr := make([]int, 1001)
	lenNums := len(nums)

	for _, arr := range nums {
		for _, num := range arr {
			mapArr[num]++
		}
	}

	for num, counter := range mapArr {
		if counter == lenNums {
			result = append(result, num)
		}
	}

	return result
}
