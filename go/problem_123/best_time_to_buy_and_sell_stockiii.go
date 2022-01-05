package problem_123

func maxProfit(prices []int) int {
	ans := 0
	onceMaxProfit := 0         // 0...i 单次交易最大收益
	minPrice := prices[0]      // 0...i 最低价格
	buyAfterOnce := -prices[0] // 0...i 完成一次交易后又买入的最大值
	for i := 1; i < len(prices); i++ {
		ans = max(ans, buyAfterOnce+prices[i])
		buyAfterOnce = max(buyAfterOnce, onceMaxProfit-prices[i])
		onceMaxProfit = max(onceMaxProfit, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
