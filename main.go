package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	boardSize    = 19
	winCondition = 5
)

func parseBoard(scanner *bufio.Scanner) [][]int {
	board := make([][]int, boardSize)
	for i := range board {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		board[i] = make([]int, boardSize)
		for j := range line {
			board[i][j], _ = strconv.Atoi(line[j])
		}
	}
	return board
}

func checkHorizontal(board [][]int) (int, int, int) {
	for i := 0; i < boardSize; i++ {
		count := 1
		for j := 1; j < boardSize; j++ {
			if board[i][j] == board[i][j-1] && board[i][j] != 0 {
				count++
			} else {
				count = 1
			}
			if count == winCondition {
				return board[i][j], i + 1, j - 4 + 1
			}
		}
	}
	return 0, 0, 0
}

func checkVertical(board [][]int) (int, int, int) {
	for j := 0; j < boardSize; j++ {
		count := 1
		for i := 1; i < boardSize; i++ {
			if board[i][j] == board[i-1][j] && board[i][j] != 0 {
				count++
			} else {
				count = 1
			}
			if count == winCondition {
				return board[i][j], i - 4 + 1, j + 1
			}
		}
	}
	return 0, 0, 0
}

func checkDiagonal(board [][]int) (int, int, int) {
	// Check top-left to bottom-right
	for i := 0; i <= boardSize-winCondition; i++ {
		for j := 0; j <= boardSize-winCondition; j++ {
			count := 1
			for k := 1; k < winCondition; k++ {
				if board[i+k][j+k] == board[i+k-1][j+k-1] && board[i+k][j+k] != 0 {
					count++
				} else {
					break
				}
				if count == winCondition {
					return board[i+k][j+k], i + 1, j + 1
				}
			}
		}
	}

	// Check top-right to bottom-left
	for i := 0; i <= boardSize-winCondition; i++ {
		for j := winCondition - 1; j < boardSize; j++ {
			count := 1
			for k := 1; k < winCondition; k++ {
				if board[i+k][j-k] == board[i+k-1][j-k+1] && board[i+k][j-k] != 0 {
					count++
				} else {
					break
				}
				if count == winCondition {
					return board[i+k][j-k], i + 1, j - k + 2
				}
			}
		}
	}
	return 0, 0, 0
}

func findWinner(board [][]int) (int, int, int) {
	winner, x, y := checkHorizontal(board)
	if winner != 0 {
		return winner, x, y
	}
	winner, x, y = checkVertical(board)
	if winner != 0 {
		return winner, x, y
	}
	winner, x, y = checkDiagonal(board)
	if winner != 0 {
		return winner, x, y
	}
	return 0, 0, 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	testCasesNumber, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error parsing number of test cases:", err)
		return
	}

	for i := 0; i < testCasesNumber; i++ {
		board := parseBoard(scanner)
		winner, x, y := findWinner(board)
		fmt.Println(winner)
		if winner != 0 {
			fmt.Println(x, y)
		}
	}
}
