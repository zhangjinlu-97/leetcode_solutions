package problem_150

import (
	"fmt"
	"testing"
)

func TestEvaluateReversePolishNotation(t *testing.T) {
	tokens := []string{"4", "13", "5", "/", "+"}
	ans := evalRPN(tokens)
	fmt.Println(ans)
}
