package problem_163

import (
	"fmt"
	"testing"
)

func TestMissingRanges(t *testing.T) {
	ans := findMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99)
	fmt.Println(ans)
}
