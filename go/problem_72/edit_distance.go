package problem_72

func minDistance(word1 string, word2 string) int {
	m, n := len(word1)+1, len(word2)+1
	dp := make([][]int, m) // dp[i][j]的意义为，word1前i个转换为word2前j个需要的最少操作数
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 1; i < m; i++ {
		dp[i][0] = i // word1的前i个，转换为word2的前0个；i次删除
	}
	for j := 0; j < n; j++ {
		dp[0][j] = j // word1的前0个，装欢为word2的前j个；j次插入
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i-1] == word2[j-1] { // word1的第i个与word2的第j个相同
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + 1 // 替换
			}
			if dp[i-1][j]+1 < dp[i][j] { // word1前i-1个转换为word2前j个，word1删除第i个
				dp[i][j] = dp[i-1][j] + 1
			}
			if dp[i][j-1]+1 < dp[i][j] { // word1前i个转换为word2前j-1个，word1再插入一个
				dp[i][j] = dp[i][j-1] + 1
			}
		}
	}
	return dp[m-1][n-1]
}
