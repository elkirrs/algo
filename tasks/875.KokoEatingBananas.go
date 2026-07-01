package tasks

import (
	"fmt"
)

func KokoEatingBananasRun() {
	fmt.Println("875. Koko Eating Bananas")
	fmt.Println("https://leetcode.com/problems/koko-eating-bananas/")
	fmt.Println("")

	kokoEatingBananasPrint([]int{3, 6, 7, 11}, 8)       // 4
	kokoEatingBananasPrint([]int{30, 11, 23, 4, 20}, 5) // 30
	kokoEatingBananasPrint([]int{30, 11, 23, 4, 20}, 6) // 23
}

func minEatingSpeed(piles []int, h int) int {
	left, right, res := 1, 0, 0

	for _, p := range piles {
		if p > right {
			right = p
		}
	}

	for left <= right {
		hours := 0
		mid := left + (right-left)/2

		for i := 0; i < len(piles); i++ {
			hours += (piles[i] + mid - 1) / mid
			if hours > h {
				break
			}
		}

		if hours <= h {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return res
}

func kokoEatingBananasPrint(piles []int, h int) {
	fmt.Println("Input: piles =", piles, "h =", h)
	out := minEatingSpeed(piles, h)
	fmt.Println("Output:", out)
}
