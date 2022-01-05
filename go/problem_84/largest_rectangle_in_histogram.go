package problem_84

// 使用单调栈，一次遍历，计算以每个柱子高度为高的最大矩形面积
// 过程中维护一个最大值，即为结果
// 栈中的[2]int数组，记录矩形高度、连续相等矩形的最右边界
func largestRectangleArea(heights []int) int {
	ret := 0
	st := make([][2]int, 0) // [2]int -> 0: 矩阵高度 1:连续相等矩形的最右边界
	st = append(st, [2]int{-1, -1})
	for i, v := range heights {
		for st[len(st)-1][0] > v {
			rec := st[len(st)-1]
			st = st[:len(st)-1]
			sq := rec[0] * (i - st[len(st)-1][1] - 1)
			if sq > ret {
				ret = sq
			}
		}
		if st[len(st)-1][0] != v {
			st = append(st, [2]int{v, i})
		} else {
			st[len(st)-1][1] = i
		}
	}
	for len(st) > 1 {
		rec := st[len(st)-1]
		st = st[:len(st)-1]
		sq := rec[0] * (len(heights) - st[len(st)-1][1] - 1)
		if sq > ret {
			ret = sq
		}
	}
	return ret
}
