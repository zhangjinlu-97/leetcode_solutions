package problem_134

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	init := -1
	for i := range gas {
		gas[i] = gas[i] - cost[i]
		if gas[i] >= 0 {
			init = i
		}
	}
	if init == -1 { // 找不到init点，则一定无答案
		return -1
	}
	need, rest, res := 0, 0, -1
	start, end := init, nextIndex(init, n)
	flag := true // 是否是首次
	for start != init || flag {
		if start != init && start == preIndex(end, n) { // start点已在连通区内
			break
		}
		if gas[start] < need {
			need -= gas[start]
		} else {
			rest += gas[start] - need
			need = 0
			for rest >= 0 && end != start {
				rest += gas[end]
				end = nextIndex(end, n)
			}
			if rest >= 0 {
				res = start
				break
			}
		}
		start = preIndex(start, n)
		flag = false
	}
	return res
}

func preIndex(index, n int) int {
	if index == 0 {
		return n - 1
	}
	return index - 1
}

func nextIndex(index, n int) int {
	if index == n-1 {
		return 0
	}
	return index + 1
}
