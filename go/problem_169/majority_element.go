package problem_169

// 摩尔投票法
// 维护res和cnt，遍历数组，当前数与res不同，则cnt-1，cnt==0时，res换成当前数
// 遍历结束后的res即为多数元素
func majorityElement(nums []int) int {
	res, cnt := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			res = nums[i]
			cnt = 1
			continue
		}
		if nums[i] == res {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}
