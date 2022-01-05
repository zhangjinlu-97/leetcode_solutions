package problem_239

import "container/list"

// 单调栈实现
func maxSlidingWindow(nums []int, k int) []int {
	dq := list.New()
	res := make([]int, 0, len(nums))
	for i := 0; i <= len(nums); i++ {
		if i-k >= 0 { // 窗口大小到达k
			front := dq.Front().Value.(int)
			res = append(res, front)
			if nums[i-k] == front {
				dq.Remove(dq.Front())
			}
		}
		if i != len(nums) {
			for dq.Len() > 0 && dq.Back().Value.(int) < nums[i] {
				dq.Remove(dq.Back())
			}
			dq.PushBack(nums[i])
		}
	}
	return res
}
