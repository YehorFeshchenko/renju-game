package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Direction to control checking mode
type Direction int

const (
	Horizontal Direction = iota
	Vertical
)

const (
	boardSize    = 19
	winCondition = 5
)

func parseBoard(scanner *bufio.Scanner) ([][]int, error) {
	board := make([][]int, boardSize)
	lineCounter := 0

	for lineCounter < boardSize {
		board[lineCounter] = make([]int, boardSize)
		if !scanner.Scan() {
			return nil, fmt.Errorf("unexpected end of input; expected 19 lines for a board, received %d", lineCounter)
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		columns := strings.Fields(line)
		if len(columns) != boardSize {
			return nil, fmt.Errorf("incorrect number of columns in row %d; expected 19 but got %d", lineCounter+1, len(columns))
		}

		for j, value := range columns {
			parsedValue, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("invalid integer value at row %d, column %d: %v", lineCounter+1, j+1, err)
			}
			board[lineCounter][j] = parsedValue
		}
		lineCounter++
	}

	return board, nil
}

// checkLines combines the checkHorizontal and checkVertical functions
func checkLines(board [][]int, direction Direction) (int, int, int) {
	for primary := 0; primary < boardSize; primary++ {
		count := 1
		for secondary := 1; secondary < boardSize; secondary++ {
			var currentValue, previousValue int
			switch direction {
			case Horizontal:
				currentValue = board[primary][secondary]
				previousValue = board[primary][secondary-1]
			case Vertical:
				currentValue = board[secondary][primary]
				previousValue = board[secondary-1][primary]
			default:
				panic("Direction exception")
			}

			if currentValue == previousValue && currentValue != 0 {
				count++
			} else {
				count = 1
			}

			if count == winCondition {
				switch direction {
				case Horizontal:
					return currentValue, primary + 1, secondary - 4 + 1
				case Vertical:
					return currentValue, secondary - 4 + 1, primary + 1
				default:
					panic("Direction exception")
				}
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
	winner, x, y := checkLines(board, Horizontal)
	if winner != 0 {
		return winner, x, y
	}
	winner, x, y = checkLines(board, Vertical)
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
	if !scanner.Scan() {
		fmt.Println("Failed to read the number of test cases")
		return
	}

	testCasesNumber, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error parsing number of test cases:", err)
		return
	}

	for i := 0; i < testCasesNumber; i++ {
		if i > 0 {
			if !scanner.Scan() || strings.TrimSpace(scanner.Text()) != "" {
				fmt.Printf("Formatting error before test case %d\n", i+1)
				continue
			}
		}
		board, err := parseBoard(scanner)
		if err != nil {
			fmt.Printf("Error parsing board for test case %d: %v\n", i+1, err)
			continue
		}
		winner, x, y := findWinner(board)
		fmt.Println(winner)
		if winner != 0 {
			fmt.Println(x, y)
		}
	}
}
