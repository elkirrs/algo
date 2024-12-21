package twoPointers

//nums := []int{3, 2, 2, 3}
//val := 3

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
