package problem_70

func climbStairs(n int) (ans int) {
	a, b := 1, 1
	fn := func() {
		a, b = b, a+b
		ans = a
	}
	for i := 0; i < n; i++ {
		fn()
	}
	return
}
