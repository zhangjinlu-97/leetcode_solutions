package problem_91

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		switch s[i] {
		case '0':
			dp[i] = 0
		case '1':
			dp[i] += dp[i+1]
			if i < n-1 {
				dp[i] += dp[i+2]
			}
		case '2':
			dp[i] += dp[i+1]
			if i < n-1 && s[i+1] <= '6' {
				dp[i] += dp[i+2]
			}
		default:
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}
