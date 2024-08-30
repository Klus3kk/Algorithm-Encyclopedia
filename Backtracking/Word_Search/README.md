# **Word Search**

## **Overview:**

The Word Search problem involves finding a word in a grid of letters. The word can be formed by connecting adjacent letters in any direction.

## **Python Example:**

```python
def word_search(board, word):
    def backtrack(r, c, idx):
        if idx == len(word):
            return True
        if not (0 <= r < len(board)) or not (0 <= c < len(board[0])) or board[r][c] != word[idx]:
            return False
        
        temp = board[r][c]
        board[r][c] = '#'
        result = (backtrack(r + 1, c, idx + 1) or
                  backtrack(r - 1, c, idx + 1) or
                  backtrack(r, c + 1, idx + 1) or
                  backtrack(r, c - 1, idx + 1))
        board[r][c] = temp
        return result

    for r in range(len(board)):
        for c in range(len(board[0])):
            if backtrack(r, c, 0):
                return True
    return False

# Example usage
board = [
    ['A', 'B', 'C', 'E'],
    ['S', 'F', 'C', 'S'],
    ['A', 'D', 'E', 'E']
]
word = "ABCCED"
print("Word found:", word_search(board, word))
```

## **Explanation:**
- **Backtracking:** Explore all possible paths from each cell and backtrack if the path does not match the word.
- **Boundary Check:** Ensure the search stays within the grid boundaries.

