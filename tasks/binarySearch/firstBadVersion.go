package binarySearch

func FirstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		verMid := left + (right-left)/2

		if isBadVersion(verMid) {
			right = verMid
		} else {
			left = verMid + 1
		}
	}

	return left
}

func isBadVersion(version int) bool {
	return version >= 4
}
