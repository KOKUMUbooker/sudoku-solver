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

func SolveSudoku(grid [][]rune) bool {
	var row, col int

	if !(FindUnassignedLocation(grid, &row, &col)) {
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
