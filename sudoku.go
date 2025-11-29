package main

func main() {
}

// Returns the grid as a slice of []runes with an individual []rune's index
// in the [][]rune slice represents it's row and an individual rune in []rune
// represents it's value in the grid
func ConvStrToRuneGrid(grid []string) [][]rune {
	//
	return [][]rune{} // TODO: Only for use in testing - CHANGE LATER
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
