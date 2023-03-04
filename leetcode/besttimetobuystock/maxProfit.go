package besttimetobuystock

// import "fmt"

func maxProfit(prices []int) int {
	var profit, low, diff int

	low = prices[0]
	for i := 1; i < len(prices); i++ {
		diff = prices[i] - low
		if prices[i] < low {
			low = prices[i]
		} else if diff > profit {
			profit = diff
		}
	}

	return profit
}
