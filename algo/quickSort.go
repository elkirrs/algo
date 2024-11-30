package algo

import "fmt"

// QuickSort O(log2n * n)
func QuickSort() {
	var list = []int{1, 4, 5, 8, 5, 1, 2, 7, 5, 2, 11}
	iterations := 0
	quickSortRun(list, &iterations)
	fmt.Printf("quickSort: len list: %d iterations: %d \n", len(list), iterations)
}

func quickSortRun(list []int, iterations *int) {
	if len(list) <= 1 {
		return
	}

	*iterations++

	var pivot int = list[0]
	var less []int
	var greater []int

	for _, val := range list[1:] {
		if val < pivot {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	quickSortRun(less, iterations)
	quickSortRun(greater, iterations)

	copy(list, append(append(less, pivot), greater...))
}
