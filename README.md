# Sudoku Solver (Go)

## Description

This project is a **Sudoku solver written in Go**.

The program takes a Sudoku grid as command-line arguments and resolves it if:

* The input is valid
* The Sudoku has **exactly one solution**

If the input is invalid or the Sudoku has no solution or multiple solutions, the program prints `Error`.

---

## Input Format

* The program expects **exactly 9 arguments**
* Each argument must be a string of **9 characters**
* Allowed characters:

  * Digits `1`–`9` → fixed numbers
  * Dot `.` → empty cell

Each argument represents **one row** of the Sudoku grid.

---

## Usage

```bash
go run . <row1> <row2> <row3> <row4> <row5> <row6> <row7> <row8> <row9>
```

---

## Example (Valid Sudoku)

```bash
$ go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
```

**Output:**

```text
3 9 6 2 4 5 7 8 1$
1 7 8 3 6 9 5 2 4$
5 2 4 8 1 7 3 9 6$
2 8 7 9 5 1 6 4 3$
9 3 1 4 8 6 2 7 5$
4 6 5 7 2 3 9 1 8$
7 1 2 6 3 8 4 5 9$
6 5 9 1 7 4 8 3 2$
8 4 3 5 9 2 1 6 7$
$
```

---

## Invalid Input Examples

### Wrong number of arguments

```bash
$ go run . 1 2 3 4 | cat -e
Error$
```

### No arguments

```bash
$ go run . | cat -e
Error$
```

### Invalid Sudoku (conflicting values)

```bash
$ go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
Error$
```

---

##  Error Handling

The program prints `Error` if:

* The number of arguments is not 9
* Any argument is not exactly 9 characters long
* Invalid characters are present
* The Sudoku is unsolvable
* The Sudoku has **more than one solution**

---

## Implementation Details

* Language: **Go**
* Solving method: **Backtracking**
* Ensures **Sudoku rules**:

  * Unique numbers (1–9) per row
  * Unique numbers per column
  * Unique numbers per 3×3 subgrid
* Only prints a solution if it is **unique**

---

## Project Structure

```
.
├── main.go
├── go.mod
└── README.md
```

---

## Testing

You can test different Sudoku configurations by modifying the input arguments.

To visualize line endings clearly (as required in some evaluations), use:

```bash
| cat -e
```

---

## Conclusion

This program validates, solves, and outputs a Sudoku grid according to strict rules, ensuring correctness and uniqueness of the solution. It follows clean input validation and standard Go practices.

