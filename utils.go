package main

// Takes a pointer to row and col & checks in grid for the next position where
// '.' is located & uses values it finds to update row & col values
// If it find a slot containing '.', it assigns row & col to those values & returns true
// else it returns false
func FindUnassignedLocation(grid [][]rune, row *int, col *int) bool {
	//
	return true // TODO: Only for use in testing - CHANGE LATER
}

// Checks whether the num argument passed against sudoku rules
// (num to be unique per row, column & also per box)
// If num is valid return true otherwise return false
func NoConflicts(grid [][]rune, row int, col int, num rune) bool {
	// Check if placement is valid along the column axis

	// Check if placement is valid along the row axis
	for k := 0; k < 9; k++ { // loop from 0 to 8
		if grid[row][k] == num { // check if current value in 'k' equals 'num'
			return false // if condition met return false, since it breaks sudoku rules
		}
	}

	// Check if placement is valid along the box

	return true // TODO: Only for use in testing - CHANGE LATER
}
