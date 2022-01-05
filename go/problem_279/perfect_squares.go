package problem_279

import (
	"math"
)

//四平方和定理：
//1. 任何数不会有超过4个完全平方数构成
//2. 任何数消掉4的因子结论不变
func numSquares(n int) int {
	for n%4 == 0 {
		n /= 4
	}
	if n%8 == 7 {
		return 4
	}
	for a := 0; a*a <= n; a++ {
		b := int(math.Sqrt(float64(n - a*a)))
		if a*a+b*b == n {
			if a > 0 && b > 0 {
				return 2
			}
			return 1
		}
	}
	return 3
}

func try(n int) int {
	nums := getNums(n)
	ans := math.MaxInt32
	var proc func(i, sum, cnt int)
	proc = func(i, sum, cnt int) {
		if sum == n {
			if cnt < ans {
				ans = cnt
			}
			return
		}
		if i == len(nums) || sum > n {
			return
		}
		proc(i, sum+nums[i], cnt+1)
		proc(i+1, sum, cnt)
	}
	proc(0, 0, 0)
	return ans
}

func getNums(n int) (nums []int) {
	for i := 1; ; i++ {
		cur := i * i
		if cur > n {
			break
		}
		nums = append(nums, cur)
	}
	return
}
