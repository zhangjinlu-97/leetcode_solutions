package problem_79

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			visited[r][c] = true
			if board[r][c] == word[0] && process(board, &word, visited, 0, r, c) {
				return true
			}
			visited[r][c] = false
		}
	}
	return false
}

func process(board [][]byte, word *string, visited [][]bool, i, r, c int) bool {
	if board[r][c] != (*word)[i] {
		return false
	}
	if i == len(*word)-1 {
		return true
	}
	ans := false
	if r > 0 && !visited[r-1][c] {
		visited[r-1][c] = true
		if process(board, word, visited, i+1, r-1, c) {
			ans = true
		}
		visited[r-1][c] = false
	}
	if r < len(board)-1 && !visited[r+1][c] {
		visited[r+1][c] = true
		if process(board, word, visited, i+1, r+1, c) {
			ans = true
		}
		visited[r+1][c] = false
	}
	if c > 0 && !visited[r][c-1] {
		visited[r][c-1] = true
		if process(board, word, visited, i+1, r, c-1) {
			ans = true
		}
		visited[r][c-1] = false
	}
	if c < len(board[0])-1 && !visited[r][c+1] {
		visited[r][c+1] = true
		if process(board, word, visited, i+1, r, c+1) {
			ans = true
		}
		visited[r][c+1] = false
	}
	return ans
}
