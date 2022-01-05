package problem_54

// 把矩阵看作一层一层的方框，维护每个方框的左上角和右下角
// 边界case: 最中心的那个方框是一条横线和竖线的情况
func spiralOrder(matrix [][]int) []int {
	tr, tc, br, bc := 0, 0, len(matrix)-1, len(matrix[0])-1
	ret := make([]int, 0, len(matrix)*len(matrix[0]))
	for tr <= br && tc <= bc {
		process(matrix, &ret, tr, tc, br, bc)
		tr++
		tc++
		br--
		bc--
	}
	return ret
}

func process(matrix [][]int, ret *[]int, tr, tc, br, bc int) {
	if tr == br {
		for c := tc; c <= bc; c++ {
			*ret = append(*ret, matrix[tr][c])
		}
	} else if tc == bc {
		for r := tr; r <= br; r++ {
			*ret = append(*ret, matrix[r][tc])
		}
	} else {
		for c := tc; c < bc; c++ {
			*ret = append(*ret, matrix[tr][c])
		}
		for r := tr; r < br; r++ {
			*ret = append(*ret, matrix[r][bc])
		}
		for c := bc; c > tc; c-- {
			*ret = append(*ret, matrix[br][c])
		}
		for r := br; r > tr; r-- {
			*ret = append(*ret, matrix[r][tc])
		}
	}
}
