package problem_15

import (
	"sort"
)

// 关键在于去重，去重的情况如下：
// 1. 如果nums[i]大于0，由于是升序排序，所以后面的和不可能小于0，直接退出循环
// 2. 如果i和上一个值相等，直接进入下一个位置，因为这种情况一定被覆盖到了
// 3. 在l~r的范围缩小的过程中，当l不是i+1且和上一个值相等时，直接跳过，因为答案是重复的
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	ret := make([][]int, 0, len(nums))
	for i, num := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && num == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		target := -num
		for l < r {
			if l != i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			sum := nums[l] + nums[r]
			if sum > target {
				r--
				continue
			} else if sum == target {
				ret = append(ret, []int{num, nums[l], nums[r]})
			}
			l++
		}
	}
	return ret
}
