// Pattern: Sliding Window (degenerate window = running minimum).
//
// Problem (LeetCode 121, Best Time to Buy and Sell Stock): given daily
// prices, return the maximum profit from one buy followed by one sell,
// or 0 if no profit is possible.
//
// Idea: the best sell day pairs with the cheapest buy day before it.
// Track the running minimum price and the best difference seen.
// Time O(n), space O(1).
package slidingwindow

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice, best := prices[0], 0
	for _, p := range prices[1:] {
		if p < minPrice {
			minPrice = p
		}
		if p-minPrice > best {
			best = p - minPrice
		}
	}
	return best
}
