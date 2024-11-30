package binarySearch

func MySqrt(x int) int {
	left, right := 0, x

	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid
		if square == x {
			return mid
		} else if square < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
