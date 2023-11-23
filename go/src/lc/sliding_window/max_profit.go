package lc

import "math"

// maxProfit (121 | Easy)
func maxProfit(prices []int) int {
	minPrice, maxProfit := math.MaxInt, 0
	for _, p := range prices {
		minPrice = min(minPrice, p)
		maxProfit = max(maxProfit, p-minPrice)
	}
	return maxProfit
}
