package problem_150

import "strconv"

func evalRPN(tokens []string) int {
	st := &Stack{}
	for _, str := range tokens {
		num, err := strconv.Atoi(str)
		if err == nil {
			st.Push(num)
		} else {
			b := st.Pop()
			a := st.Pop()
			st.Push(calc(str, a, b))
		}
	}
	return st.Pop()
}

func calc(opt string, a, b int) int {
	ans := 0
	if opt == "+" {
		ans = a + b
	} else if opt == "-" {
		ans = a - b
	} else if opt == "*" {
		ans = a * b
	} else if opt == "/" {
		ans = a / b
	}
	return ans
}

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}
