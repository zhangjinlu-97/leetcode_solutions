package problem_76

import "math"

// sliding window
func minWindow(s string, t string) string {
	sMap, tMap := make([]int, 128), make([]int, 128)
	for i := range t {
		tMap[t[i]]++
	}
	minStr, minLen := "", math.MaxInt64
	l, r := 0, 0 // r is the first wrong index
	for l < len(s) {
		for r < len(s) && !checkStr(s, sMap, tMap) {
			sMap[s[r]]++
			r++
		}
		if checkStr(s, sMap, tMap) && r-l < minLen {
			minStr = s[l:r]
			minLen = r - l
		}
		sMap[s[l]]--
		l++
	}
	return minStr
}

// O(1) Complexity
func checkStr(str string, sMap []int, tMap []int) bool {
	for i := range tMap {
		if sMap[i] < tMap[i] {
			return false
		}
	}
	return true
}
