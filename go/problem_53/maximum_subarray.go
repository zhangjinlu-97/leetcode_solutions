package problem_53

import "math"

func maxSubArray(nums []int) int {
	ret, preSum := math.MinInt32, 0
	for i := range nums {
		preSum += nums[i] // 以当前位置结尾的最大子列和
		if preSum > ret {
			ret = preSum
		}
		if preSum < 0 {
			preSum = 0 // 若pre<0则以i+1结尾的最大子序列肯定不加pre
		}
	}
	return ret
}
