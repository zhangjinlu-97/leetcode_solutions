package problem_239

import (
	"fmt"
	"testing"
)

func TestSlidingWindowMaximum(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := maxSlidingWindow(nums, k)
	fmt.Println(res)
}
