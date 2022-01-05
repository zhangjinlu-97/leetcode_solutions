package problem_139

// 前缀树优化动态规划
func wordBreak(s string, wordDict []string) bool {
	root := initTrie(wordDict)
	n := len(s)
	str := []byte(s)
	dp := make([]bool, n+1)
	dp[n] = true
	for i := n - 1; i >= 0; i-- {
		cur := root
		for j := i; j < n; j++ {
			cur = cur.nextMap[str[j]-'a']
			if cur == nil { // 如果后序没有路径了，直接break
				break
			}
			if cur.end && dp[j+1] { // str[i:j+1]存在则加上
				dp[i] = true
			}
		}
	}
	return dp[0]
}

type Node struct {
	end     bool
	nextMap []*Node
}

func initTrie(wordDict []string) *Node {
	root := &Node{nextMap: make([]*Node, 26)}
	for i := range wordDict {
		str := []byte(wordDict[i])
		cur := root
		for _, b := range str {
			path := b - 'a'
			if cur.nextMap[path] == nil {
				cur.nextMap[path] = &Node{nextMap: make([]*Node, 26)}
			}
			cur = cur.nextMap[path]
		}
		cur.end = true
	}
	return root
}

// 递归解法
func wordBreak1(s string, wordDict []string) bool {
	wd := make(map[string]struct{})
	for _, str := range wordDict {
		wd[str] = struct{}{}
	}
	var prod func(i int) bool
	prod = func(i int) bool {
		if i == len(s) {
			return true
		}
		ans := false
		for j := i + 1; j <= len(s); j++ {
			if _, ok := wd[s[i:j]]; ok && prod(j) {
				ans = true
			}
		}
		return ans
	}
	return prod(0)
}

// 动态规划
func wordBreak2(s string, wordDict []string) bool {
	wd := make(map[string]struct{})
	for _, str := range wordDict {
		wd[str] = struct{}{}
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[n] = true
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j <= n; j++ {
			if _, ok := wd[s[i:j]]; ok && dp[j] {
				dp[i] = true
			}
		}
	}
	return dp[0]
}
