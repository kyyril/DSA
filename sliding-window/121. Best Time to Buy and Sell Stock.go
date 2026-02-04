package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}