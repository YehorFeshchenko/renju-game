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

func parseBoard(scanner *bufio.Scanner) ([][]int, error) {
	board := make([][]int, boardSize)
	linesRead := 0

	for linesRead < boardSize {
		board[linesRead] = make([]int, boardSize)
		if !scanner.Scan() {
			return nil, fmt.Errorf("unexpected end of input; expected 19 lines for a board, received %d", linesRead)
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		columns := strings.Fields(line)
		if len(columns) != boardSize {
			return nil, fmt.Errorf("incorrect number of columns in row %d; expected 19 but got %d", linesRead+1, len(columns))
		}

		for j, value := range columns {
			parsedValue, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("invalid integer value at row %d, column %d: %v", linesRead+1, j+1, err)
			}
			board[linesRead][j] = parsedValue
		}
		linesRead++
	}

	return board, nil
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

func validateTestCases(scanner *bufio.Scanner, numCases int) error {
	expectedRows := boardSize * numCases
	actualRows := 0
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		lineCount++

		if trimmedLine == "" || len(strings.Fields(trimmedLine)) == 0 { // Skip empty or space-only lines
			fmt.Printf("Skipping empty or invalid line at physical line %d\n", lineCount)
			continue
		}

		columns := strings.Fields(trimmedLine)
		if len(columns) != boardSize {
			fmt.Printf("Skipping line with incorrect format at physical line %d: found %d columns\n", lineCount, len(columns))
			continue
		}

		actualRows++
		fmt.Printf("Processed valid line %d (physical line %d): %s\n", actualRows, lineCount, line)
	}

	if actualRows != expectedRows {
		return fmt.Errorf("validation failed: expected %d rows, got %d rows", expectedRows, actualRows)
	}
	return nil
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
