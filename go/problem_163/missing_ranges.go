package problem_163

import (
	"fmt"
	"strconv"
)

func findMissingRanges(nums []int, lower int, upper int) []string {
	if len(nums) == 0 {
		return []string{getRange(lower, upper)}
	}
	ans := make([]string, 0)
	cur := lower
	for _, num := range nums {
		if num != cur {
			ans = append(ans, getRange(cur, num-1))
			cur = num + 1
			continue
		}
		cur++
	}
	if last := nums[len(nums)-1]; last != upper {
		ans = append(ans, getRange(last+1, upper))
	}
	return ans
}

func getRange(lower, upper int) string {
	if upper == lower {
		return strconv.Itoa(lower)
	}
	return fmt.Sprintf("%d->%d", lower, upper)
}
