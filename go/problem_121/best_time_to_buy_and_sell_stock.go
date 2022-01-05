package problem_121

// 遍历数组，维护之前的最小价格 和 当前的最大收益即可
func maxProfit(prices []int) int {
	preMin := prices[0]
	ret := 0
	for i := 1; i < len(prices); i++ {
		ret = max(ret, prices[i]-preMin)
		preMin = min(preMin, prices[i])
	}
	return ret
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
