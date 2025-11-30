package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1:]
	grid := ConvStrToRuneGrid(input)
	if grid == nil {
		fmt.Println("Error")
		return
	}

	SolveSudoku(grid)

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func ConvStrToRuneGrid(grid []string) [][]rune {
	if len(grid) != 9 {
		return nil
	}

	gridRune := make([][]rune, 9)
	for i, row := range grid {
		if len(row) != 9 {
			return nil
		}
		gridRune[i] = []rune(row)
	}

	return gridRune
}

func FindUnassignedLocation(grid [][]rune, row *int, col *int) bool {
	dotFound := false

	for rowI, rowSlice := range grid {
		for colI, r := range rowSlice {
			if r == '.' {
				dotFound = true
				*row = rowI
				*col = colI
				break
			}
		}

		if dotFound {
			break
		}
	}

	return dotFound
}

func NoConflicts(grid [][]rune, row int, col int, num rune) bool {
	// Check if placement is valid along the column axis
	for r := 0; r < 9; r++ {
		if grid[r][col] == num {
			return false
		}
	}

	// Check if placement is valid along the row axis
	for k := 0; k < 9; k++ {
		if grid[row][k] == num {
			return false
		}
	}

	// Check if placement is valid along the box
	boxRowStart := (row / 3) * 3
	boxColStart := (col / 3) * 3

	// Loop through the 3×3 box
	for r := boxRowStart; r < boxRowStart+3; r++ {
		for c := boxColStart; c < boxColStart+3; c++ {
			if grid[r][c] == num {
				return false
			}
		}
	}

	return true
}

func SolveSudoku(grid [][]rune) bool {
	var row, col int

	// Base condition
	openSlotFound := FindUnassignedLocation(grid, &row, &col)
	if openSlotFound {
		return true
	}

	for num := '1'; num <= '9'; num++ {
		if NoConflicts(grid, row, col, num) {
			grid[row][col] = num
			if SolveSudoku(grid) {
				return true
			}
			grid[row][col] = '.'
		}
	}

	return false
}
