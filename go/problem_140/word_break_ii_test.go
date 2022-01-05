package problem_140

import (
	"fmt"
	"testing"
)

func TestWordBreakII(t *testing.T) {
	str := "catsanddog"
	wd := []string{"cat", "cats", "and", "sand", "dog"}
	ans := wordBreak(str, wd)
	fmt.Println(ans)
}
