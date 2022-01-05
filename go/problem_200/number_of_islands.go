package problem_200

//遍历矩阵，遇到 1 ，则ret++，并递归的将这一片陆地全部标记为#，避免后面重复计数
func numIslands(grid [][]byte) int {
	ret := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ret++
				process(grid, i, j)
			}
		}
	}
	return ret
}

func process(grid [][]byte, row, col int) {
	if grid[row][col] != '1' {
		return
	}
	grid[row][col] = '#'
	if row > 0 {
		process(grid, row-1, col)
	}
	if row < len(grid)-1 {
		process(grid, row+1, col)
	}
	if col > 0 {
		process(grid, row, col-1)
	}
	if col < len(grid[0])-1 {
		process(grid, row, col+1)
	}
}
