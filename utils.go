package main

// Takes a pointer to row and col & checks in grid for the next position where
// '.' is located & uses values it finds to update row & col values
// If it find a slot containing '.', it assigns row & col to those values & returns true
// else it returns false
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

// Checks whether the num argument passed against sudoku rules
// (num to be unique per row, column & also per box)
// If num is valid return true otherwise return false
func NoConflicts(grid [][]rune, row int, col int, num rune) bool {
	// Check if placement is valid along the column axis
	// Check if placement is valid along the row axis

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
