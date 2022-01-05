package problem_279

import (
	"fmt"
	"testing"
)

func TestPerfectSquares(t *testing.T) {
	for i := 1; i < 1000; i++ {
		ans := try(i)
		fmt.Printf("%d: %d\n", i, ans)
	}
}
