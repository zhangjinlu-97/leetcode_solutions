package problem_33

// 直接二分，如果nums[l] <= nums[m]说明，m在左半段，l~m是升序的，可以判断target是否在l~m之间，如果不在，则排除掉l~m，直接在m+1~r上二分。
// 注意一定是 <= 因为l和m可能相等
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return m
		}
		if nums[l] <= nums[m] {
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}
