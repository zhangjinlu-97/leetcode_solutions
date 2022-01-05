package problem_42

// lmax, rmax中较小的那个就是l或r上水量的瓶颈
// 此时，计算max(0, lmax-height[l] 或 rmax-height[r])即可求出此柱子上的最大水量
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	water, lmax, rmax := 0, height[0], height[len(height)-1]
	l, r := 1, len(height)-2
	for l <= r {
		if lmax <= rmax {
			water += max(0, lmax-height[l])
			lmax = max(lmax, height[l])
			l++
		} else {
			water += max(0, rmax-height[r])
			rmax = max(rmax, height[r])
			r--
		}
	}
	return water
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
