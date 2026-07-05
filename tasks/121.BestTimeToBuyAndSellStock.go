package tasks

import (
	"fmt"
)

func BestTimeToBuyAndSellStockRun() {
	fmt.Printf("121. Best Time to Buy and Sell Stock")
	fmt.Println("https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/")
	fmt.Printf("")

	bestTimeToBuyAndSellStockPrint([]int{7, 1, 5, 3, 6, 4})
	bestTimeToBuyAndSellStockPrint([]int{7, 6, 4, 3, 1})
}

func maxProfit(prices []int) int {
	l := 0
	tmp := prices[0]

	for r := 1; r < len(prices); r++ {
		if tmp > prices[r] {
			tmp = prices[r]
		} else {
			c := prices[r] - tmp
			if c > l {
				l = c
			} 
		}
	}

	return l
}

func bestTimeToBuyAndSellStockPrint(in []int) {
	fmt.Println("Input:", in)
	out := maxProfit(in)
	fmt.Println("Output:", out)
}
