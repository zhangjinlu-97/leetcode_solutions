package problem_3

//枚举以每个位置结尾的情况，取最大
//1. i位置字符上次出现的下标a
//2. i-1位置字符推不到的第一个位置b
//i - max(a, b)即为以i结尾的最长无重复子串的长度
func lengthOfLongestSubstring(s string) int {
	idxMap := make([]int, 256)
	for i := range idxMap {
		idxMap[i] = -1
	}
	pre := -1
	ret := 0
	for i, c := range s {
		pre = max(pre, idxMap[c])
		ret = max(ret, i-pre)
		idxMap[c] = i
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
