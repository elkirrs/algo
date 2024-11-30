package algo

import "fmt"

func BubbleSort() {
	list := []int{1, 4, 5, 8, 5, 1, 2, 7, 5, 2, 11}
	bubbleSortRun(&list)
}

// bubbleSortRun O(n*2)
func bubbleSortRun(list *[]int) {
	slice := *list
	var listLen int = len(slice) - 1
	var tmpVal int
	var iterations int = 0

	for i := 1; i < listLen; i++ {
		for j := 1; j < listLen; j++ {
			if slice[j+1] < slice[j] {
				tmpVal = slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = tmpVal
			}
			iterations++
		}
	}
	fmt.Printf("bubbleSort:  len list: %d iterations: %d \n", len(slice), iterations)
}
