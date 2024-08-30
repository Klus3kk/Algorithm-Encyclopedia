# **Matrix DP**

## **Overview:**

Matrix DP problems involve optimization over a 2D matrix and typically involve computing values for each cell based on previous cells.

## **Python Example:**

- **Longest Common Subsequence:**

```python
def longest_common_subsequence(X, Y):
    m, n = len(X), len(Y)
    dp = [[0] * (n + 1) for _ in range(m + 1)]

    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if X[i - 1] == Y[j - 1]:
                dp[i][j] = dp[i - 1][j - 1] + 1
            else:
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])

    return dp[m][n]

# Example usage
X = "ABCBDAB"
Y = "BDCAB"
print("Length of Longest Common Subsequence:", longest_common_subsequence(X, Y))
```

- **Longest Increasing Path in Matrix:**

```python
def longest_increasing_path(matrix):
    if not matrix:
        return 0

    def dfs(x, y):
        if memo[x][y] != -1:
            return memo[x][y]

        length = 1
        for dx, dy in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            nx, ny = x + dx, y + dy
            if 0 <= nx < m and 0 <= ny < n and matrix[nx][ny] > matrix[x][y]:
                length = max(length, 1 + dfs(nx, ny))

        memo[x][y] = length
        return length

    m, n = len(matrix), len(matrix[0])
    memo = [[-1] * n for _ in range(m)]

    return max(dfs(x, y) for x in range(m) for y in range(n))

# Example usage
matrix = [
    [9, 9, 4],
    [6, 6, 8],
    [2, 1, 1]
]
print("Length of Longest Increasing Path:", longest_increasing_path(matrix))
```

## **Explanation:**
- **Longest Common Subsequence:** Finds the longest subsequence that is common between two sequences.
- **Longest Increasing Path:** Finds the longest path in a matrix where each cell value is greater than the previous cell value.

