package problem_125

import "strings"

func isPalindrome(s string) bool {
	var sb strings.Builder
	for i := range s {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			if c >= 'A' && c <= 'Z' {
				c = 'a' + c - 'A'
			}
			sb.WriteByte(c)
		}
	}
	s = sb.String()
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
