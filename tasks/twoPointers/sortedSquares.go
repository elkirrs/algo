package twoPointers

//nums = []int{-4,-1,0,3,10}
//Output[0,1,9,16,100]

func sortedSquares(nums []int) []int {
	n := len(nums)
	slice := make([]int, len(nums))
	l, r := 0, n-1
	pos := n - 1

	for l <= r {
		lSquare := nums[l] * nums[l]
		rSquare := nums[r] * nums[r]

		if lSquare > rSquare {
			slice[pos] = lSquare
			l++
		} else {
			slice[pos] = rSquare
			r--
		}

		pos--
	}

	return slice
}
