package problem_122

func maxProfit(prices []int) int {
	ans := 0
	for i := range prices {
		if i > 0 && prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
