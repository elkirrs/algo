package algo

import "fmt"

func LineSearch() {
	var list = []int{1, 4, 5, 8, 5, 1, 2, 7, 5, 2, 11}
	var target = 7

	lineSearchRun(list, target)
}

// lineSearchRun O(n)
func lineSearchRun(list []int, target int) int {
	var i = -1
	var iterations = 0

	for idx, val := range list {
		iterations++
		if val == target {
			i = idx
			break
		}
	}

	fmt.Printf("lineSearch: value: %d index: %d iterations: %d \n", target, i, iterations)

	return i
}
