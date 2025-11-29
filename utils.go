package main

// Takes a pointer to row and col & checks in grid for the next position where
// '.' is located & uses values it finds to update row & col values
// If it find a slot containing '.', it assigns row & col to those values & returns true
// else it returns false
func FindUnassignedLocation(grid [][]rune, row *int, col *int) bool {
}

// Checks whether the num argument passed against sudoku rules
// (num to be unique per row, column & also per box)
// If num is valid return true otherwise return false
func NoConflicts(grid [][]rune, row int, col int, num rune) bool {
}
