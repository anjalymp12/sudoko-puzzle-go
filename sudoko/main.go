package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Holds starting puzzle
	puzz := [][]int{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}

	displayBoard(puzz)
	fmt.Println()
	solvePuzz(puzz)
	displayBoard(puzz)

}

func isNumValid(puzz [][]int, guess, row, colum int) bool {
	for i := range puzz {
		if puzz[row][i] == guess && colum != i {
			return false
		}
	}
	for i := range puzz {
		if puzz[i][colum] == guess && row != i {
			return false
		}
	}
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			if puzz[row-row%3+k][colum-colum%3+l] == guess && (row-row%3+k != row) || (colum-colum%3+l != colum) {
				return true
			}
		}

	}
	return true
}
func getEmptySpaces(puzz [][]int) (int, int) {
	puzz_cols := len(puzz[0])
	for i := 0; i < len(puzz); i++ {
		for j := 0; j < (puzz_cols); j++ {
			if puzz[i][j] == 0 {
				return i, j
			}
		}

	}
	return -1, -1

}

// Outputs the current board to screen
func displayBoard(puzz [][]int) {
	puzz_cols := len(puzz[0])
	// Cycles through rows in the puzz
	for i := 0; i < len(puzz); i++ {
		// Cycle through columns in puzz
		for j := 0; j < puzz_cols; j++ {
			fmt.Print(strconv.Itoa(puzz[i][j]) + " ")
		}
		fmt.Print("\n")
	}
}

func solvePuzz(puzz [][]int) bool {
	row, colum := getEmptySpaces(puzz)
	if row == -1 {
		return true
	} 
	for i := 1; i <= 9; i++ {
		if isNumValid(puzz, i, row, colum) {
			puzz[row][colum] = i
			if solvePuzz(puzz) {
				return true
			}
			puzz[row][colum] = 0
		}
	}
	return false
}
