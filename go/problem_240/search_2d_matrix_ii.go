package problem_240

//解法1
// 特殊遍历，从右上角开始遍历，大了往左，小了往右
// m行数， n列数，时间复杂度: O(m + n)
func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		} else {
			return true
		}
	}
	return false
}

// 解法2
// 逐 行/列 二分
// m行数， n列数，时间复杂度: O(min(m,n) * log(max(m, n)))
func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) > len(matrix[0]) { // 行数大于列数
		for i := range matrix[0] {
			if colSearch(matrix, target, i) {
				return true
			}
		}
		return false
	}
	for i := range matrix {
		if rowSearch(matrix, target, i) {
			return true
		}
	}
	return false
}

func rowSearch(matrix [][]int, target, i int) bool {
	l, r := 0, len(matrix[i])-1
	for l <= r {
		m := l + (r-l)>>1
		if matrix[i][m] < target {
			l = m + 1
		} else if matrix[i][m] > target {
			r = m - 1
		} else {
			return true
		}
	}
	return false
}

func colSearch(matrix [][]int, target, i int) bool {
	l, r := 0, len(matrix)-1
	for l <= r {
		m := l + (r-l)>>1
		if matrix[m][i] < target {
			l = m + 1
		} else if matrix[m][i] > target {
			r = m - 1
		} else {
			return true
		}
	}
	return false
}
