package binarySearch

//mat = [][]int{
//{10, 20, 15},
//{21, 30, 14},
//{7, 16, 32},
//}

func findPeakGrid(mat [][]int) []int {
	result := []int{-1, -1}

	if len(mat) == 0 && len(mat[0]) == 0 {
		return result
	}

	left, right := 0, len(mat[0])-1

	for left <= right {
		mid := left + (right-left)/2

		max := 0
		for i := 0; i < len(mat); i++ {
			if mat[i][mid] > mat[max][mid] {
				max = i
			}
		}

		if (mid == 0 || mat[max][mid] > mat[max][mid-1]) && (mid == len(mat[0])-1 || mat[max][mid] > mat[max][mid+1]) {
			result[0] = max
			result[1] = mid
			return result
		} else if mid > 0 && mat[max][mid] < mat[max][mid-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}
