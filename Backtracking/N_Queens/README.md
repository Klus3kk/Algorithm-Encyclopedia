# **N-Queens Problem**

## **Overview:**

The N-Queens problem involves placing N queens on an N×N chessboard such that no two queens threaten each other. This problem is a classic example of constraint satisfaction and backtracking.

## **Python Example:**

```python
def solve_n_queens(n):
    def is_valid(board, row, col):
        for i in range(row):
            if board[i] == col or \
               board[i] - i == col - row or \
               board[i] + i == col + row:
                return False
        return True

    def solve(board, row):
        if row == n:
            solutions.append(board[:])
            return
        for col in range(n):
            if is_valid(board, row, col):
                board[row] = col
                solve(board, row + 1)
                board[row] = -1
    
    solutions = []
    board = [-1] * n
    solve(board, 0)
    return solutions

# Example usage
n = 4
solutions = solve_n_queens(n)
print("Number of solutions for", n, "Queens:", len(solutions))
print("Solutions:", solutions)
```

## **Explanation:**
- **Validity Check:** Ensure no two queens threaten each other.
- **Backtracking:** Place queens row by row and backtrack when a placement does not lead to a solution.

