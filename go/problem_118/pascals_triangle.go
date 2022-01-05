package problem_118

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := range ret {
		ret[i] = make([]int, i+1)
		ret[i][0] = 1
		ret[i][i] = 1
		for j := 1; j < i; j++ {
			ret[i][j] = ret[i-1][j] + ret[i-1][j-1]
		}
	}
	return ret
}
