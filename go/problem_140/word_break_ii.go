package problem_140

func wordBreak(s string, wordDict []string) []string {
	root := newTrie(wordDict)
	n := len(s)
	dp := make([]bool, n+1)
	dp[n] = true
	str := []byte(s)
	for i := n - 1; i >= 0; i-- {
		cur := root
		for j := i; j < n; j++ {
			cur = cur.nexts[str[j]-'a']
			if cur == nil {
				break
			}
			if cur.end && dp[j+1] {
				dp[i] = true
			}
		}
	}
	ans := make([]string, 0)
	path := make([]byte, 0)
	var proc func(i int)
	proc = func(i int) {
		if i == len(str) {
			ans = append(ans, string(path))
			return
		}
		cur := root
		for j := i; j < n; j++ {
			index := str[j] - 'a'
			if cur.nexts[index] == nil {
				break
			}
			cur = cur.nexts[index]
			if cur.end && dp[j+1] { // 前一段字符串再字典中 且 接下来的字符串能被分割
				flag := false
				if len(path) != 0 {
					path = append(path, ' ')
					flag = true
				}
				path = append(path, cur.path...)
				proc(j + 1)
				path = path[:len(path)-len(cur.path)]
				if flag {
					path = path[:len(path)-1]
				}
			}
		}
	}
	proc(0)
	return ans
}

type node struct {
	end   bool
	path  []byte
	nexts []*node
}

func newTrie(wordDict []string) *node {
	root := &node{nexts: make([]*node, 26)}
	for _, s := range wordDict {
		str := []byte(s)
		cur := root
		for _, b := range str {
			index := b - 'a'
			if cur.nexts[index] == nil {
				cur.nexts[index] = &node{nexts: make([]*node, 26)}
			}
			cur = cur.nexts[index]
		}
		cur.path = str
		cur.end = true
	}
	return root
}
