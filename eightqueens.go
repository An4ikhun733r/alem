package piscine

import "github.com/01-edu/z01"

const N = 8

func solveQueens() {
	board := make([]int, N)
	solve(board, 0)
}

func solve(board []int, col int) {
	if col == N {
		printBoard(board)
		return
	}

	for row := 0; row < N; row++ {
		if isSafe(board, row, col) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

func isSafe(board []int, row, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row || board[i]-row == i-col || board[i]-row == col-i {
			return false
		}
	}
	return true
}

func printBoard(board []int) {
	for _, row := range board {
		for i := 0; i < N; i++ {
			if i == row {
				z01.PrintRune(rune(i + 1))
			} else {
				z01.PrintRune('.')
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	solveQueens()
}
