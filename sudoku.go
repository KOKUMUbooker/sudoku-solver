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

	// Ensure default grid is valid before solving
	for rowI, rowRunes := range grid {
		for colI, r := range rowRunes {
			if r == '.' {
				continue
			}
			num := grid[rowI][colI]
			count := CheckGridIsValidBeforeSolving(grid, rowI, colI, num)
			if count > 0 {
				fmt.Println("Error")
				return
			}
		}
	}

	sudokuSolved := SolveSudoku(grid)
	if sudokuSolved == false {
		fmt.Println("Error")
		return
	}

	for _, row := range grid {
		for i, r := range row {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(string(r))
		}
		fmt.Println()
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
	for rowI, rowSlice := range grid {
		for colI, r := range rowSlice {
			if r == '.' {
				*row = rowI
				*col = colI
				return true
			}
		}
	}

	return false
}

func NoConflicts(grid [][]rune, row int, col int, num rune) bool {
	for r := 0; r < 9; r++ {
		if grid[r][col] == num {
			return false
		}
	}

	for k := 0; k < 9; k++ {
		if grid[row][k] == num {
			return false
		}
	}

	boxRowStart := (row / 3) * 3
	boxColStart := (col / 3) * 3

	for r := boxRowStart; r < boxRowStart+3; r++ {
		for c := boxColStart; c < boxColStart+3; c++ {
			if grid[r][c] == num {
				return false
			}
		}
	}

	return true
}

func CheckGridIsValidBeforeSolving(grid [][]rune, row int, col int, num rune) int {
	count := 0
	for r := 0; r < 9; r++ {
		if grid[r][col] == num && r != row {
			count++
		}
	}

	for k := 0; k < 9; k++ {
		if grid[row][k] == num && k != col {
			count++
		}
	}

	boxRowStart := (row / 3) * 3
	boxColStart := (col / 3) * 3

	for r := boxRowStart; r < boxRowStart+3; r++ {
		for c := boxColStart; c < boxColStart+3; c++ {
			if grid[r][c] == num && r != row && c != col {
				count++
			}
		}
	}

	return count
}

func SolveSudoku(grid [][]rune) bool {
	var row, col int

	openSlotFound := FindUnassignedLocation(grid, &row, &col)
	if openSlotFound == false {
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
