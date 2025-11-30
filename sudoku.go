package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1:] // evrything after program name

	grid := ConvStrToRuneGrid(input)
	if grid == nil {
		fmt.Println("Error") // invalid input
		return
	}

	for _, row := range grid { // TODO: Only for use in testing - DELETE LATER
		fmt.Println(string(row))
	}
}

// Returns the grid as a slice of []runes with an individual []rune's index
// in the [][]rune slice represents it's row and an individual rune in []rune
// represents it's value in the grid
func ConvStrToRuneGrid(grid []string) [][]rune {
	if len(grid) != 9 {
		return nil
	}

	gridRune := make([][]rune, 9) // allocate 2d slice to hold the 9 rows
	for i, row := range grid {
		if len(row) != 9 { // each roe must contain 9 characters
			return nil // otherwise not valid
		}
		gridRune[i] = []rune(row) // convert string to rune slice
	}

	return gridRune // if everything is valid, return the full 9x9 grid
}

// Solves sudoku by recursion & backtracking
func SolveSudoku(grid [][]rune) bool {
	var row, col int

	// Checks for next '.' location & assigns it the values for row & col if found assigns them & returns true
	// Otherwise returns false
	if !(FindUnassignedLocation(grid, &row, &col)) { // Base condition
		return true // All '.' locations got assigned
	}

	for num := '1'; num <= '9'; num++ {
		if NoConflicts(grid, row, col, num) {
			grid[row][col] = num // Try the choic since it doesn't violate sudoku rules
			// Since it doesn't violate soduku rules be optimistic that starting from the filled position
			// onwards will hopefully yield a solved soduku
			if SolveSudoku(grid) {
				return true
			}
			grid[row][col] = '.' // Undo the choice
		}
	}

	return false
}
