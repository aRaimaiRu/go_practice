package besttimetobuystock

// import "fmt"

func maxProfit(prices []int) int {
	var profit, candidate, ihigh int

	low := prices[0]
	candidate = prices[0]
	// fmt.Println("profit, prices[i], low, candidate")
	for i := 1; i < len(prices); i++ {
		// fmt.Println(profit, prices[i], low, candidate)
		if ihigh == 0 && (prices[i]-low) < 0 {
			low = prices[i]
			candidate = low
		}
		if (prices[i] - candidate) > profit {
			profit = (prices[i] - candidate)
			ihigh = i
		}

		if low > prices[i] && candidate > prices[i] {
			candidate = prices[i]
		}

	}

	return profit
}
