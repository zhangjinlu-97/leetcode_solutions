package problem_20

// 遍历字符串，遇到左括号入栈
//遇到右括号出栈，判断是否匹配
func isValid(s string) bool {
	st := &stack{}
	str := []byte(s)
	for _, c := range str {
		if c == '(' || c == '[' || c == '{' {
			st.push(c)
			continue
		}
		if st.isEmpty() {
			return false
		}
		pre := st.pop()
		if (pre == '(' && c != ')') ||
			(pre == '[' && c != ']') ||
			(pre == '{' && c != '}') {
			return false
		}
	}
	return st.isEmpty()
}

type stack []byte

func (st *stack) push(b byte) {
	*st = append(*st, b)
}

func (st *stack) pop() byte {
	old := *st
	n := len(old)
	b := old[n-1]
	*st = old[:n-1]
	return b
}

func (st *stack) isEmpty() bool {
	return len(*st) == 0
}
