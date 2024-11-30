package algo

import "fmt"

func SelectionSort() {
	var list = []int{1, 4, 5, 8, 5, 1, 2, 7, 5, 2, 11}
	selectionSortRun(&list)
}

// selectionSortRun O(n*2)
func selectionSortRun(list *[]int) {
	slice := *list
	var indexMin int = 0
	var listLen int = len(slice) - 1
	var tmpVal int
	var iterations int = 0

	for i := 0; i < listLen; i++ {
		indexMin = i
		for j := i + 1; j < listLen; j++ {
			if slice[j] < slice[indexMin] {
				indexMin = j
			}
			iterations++
		}
		tmpVal = slice[i]
		slice[i] = slice[indexMin]
		slice[indexMin] = tmpVal
	}
	fmt.Printf("selectionSort: len list: %d iterations: %d \n", len(slice), iterations)
}
