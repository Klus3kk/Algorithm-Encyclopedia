# **Sudoku Solver**

## **Overview:**

The Sudoku Solver fills a 9×9 grid with digits so that each column, row, and 3×3 sub-grid contains all digits from 1 to 9 without repetition. It is a classic example of constraint satisfaction.

## **Python Example:**

```python
def solve_sudoku(board):
    def is_valid(board, row, col, num):
        for i in range(9):
            if board[row][i] == num or board[i][col] == num or \
               board[row - row % 3 + i // 3][col - col % 3 + i % 3] == num:
                return False
        return True

    def solve(board):
        for row in range(9):
            for col in range(9):
                if board[row][col] == 0:
                    for num in range(1, 10):
                        if is_valid(board, row, col, num):
                            board[row][col] = num
                            if solve(board):
                                return True
                            board[row][col] = 0
                    return False
        return True

    solve(board)
    return board

# Example usage
board = [
    [5, 3, 0, 0, 7, 0, 0, 0, 0],
    [6, 0, 0, 1, 9, 5, 0, 0, 0],
    [0, 9, 8, 0, 0, 0, 0, 6, 0],
    [8, 0, 0, 0, 6, 0, 0, 0, 3],
    [4, 0, 0, 8, 0, 3, 0, 0, 1],
    [7, 0, 0, 0, 2, 0, 0, 0, 6],
    [0, 6, 0, 0, 0, 0, 2, 8, 0],
    [0, 0, 0, 4, 1, 9, 0, 0, 5],
    [0, 0, 0, 0, 8, 0, 0, 7, 9]
]

solved_board = solve_sudoku(board)
print("Solved Sudoku Board:")
for row in solved_board:
    print(row)
```

## **Explanation:**
- **Validity Check:** Ensure that placing a number does not violate Sudoku rules.
- **Backtracking:** Try placing each number and backtrack if a number cannot be placed validly.

