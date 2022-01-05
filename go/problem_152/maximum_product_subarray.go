package problem_152

//如何获取以i位置结尾的子数组的最大累乘积和最小累乘积？
//三种情况：
//1. 只包含i位置
//2. i位置 * i-1位置最大
//3. i位置 * i-1位置最小
func maxProduct(nums []int) int {
	preMax, preMin, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		curMax := max(num, max(preMax*num, preMin*num))
		curMin := min(num, min(preMax*num, preMin*num))
		ans = max(ans, curMax)
		preMax = curMax
		preMin = curMin
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
