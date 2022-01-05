package problem_218

import (
	"fmt"
	"testing"
)

func TestSkipListMap(t *testing.T) {
	param := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}
	skyline := getSkyline(param)
	fmt.Println(skyline)
}
