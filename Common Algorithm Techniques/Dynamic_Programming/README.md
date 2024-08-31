# **Dynamic Programming**

## **Overview:**

Dynamic Programming (DP) is a technique for solving problems by breaking them down into simpler subproblems and storing the results of these subproblems to avoid redundant computations. It is particularly useful for optimization problems.

## **Subtopics:**

- **Knapsack Problem:** Determines the maximum value that can be carried in a knapsack of limited capacity.
- **Longest Common Subsequence (LCS):** Finds the longest subsequence common to two sequences.
- **Matrix Chain Multiplication:** Finds the most efficient way to multiply a chain of matrices.

## **Python Example (Knapsack Problem):**

```python
def knapsack(weights, values, capacity):
    n = len(weights)
    dp = [0] * (capacity + 1)

    for i in range(n):
        for w in range(capacity, weights[i] - 1, -1):
            dp[w] = max(dp[w], dp[w - weights[i]] + values[i])

    return dp[capacity]

# Example usage
weights = [1, 2, 3]
values = [60, 100, 120]
capacity = 5
print("Maximum value in knapsack:", knapsack(weights, values, capacity))
```

## **Explanation:**
- **Knapsack Problem:** Uses dynamic programming to maximize the total value of items that can be included in the knapsack without exceeding its capacity.

## **Python Example (Longest Common Subsequence):**

```python
def lcs(X, Y):
    m = len(X)
    n = len(Y)
    dp = [[0] * (n + 1) for _ in range(m + 1)]

    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if X[i - 1] == Y[j - 1]:
                dp[i][j] = dp[i - 1][j - 1] + 1
            else:
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])

    return dp[m][n]

# Example usage
X = "AGGTAB"
Y = "GXTXAYB"
print("Length of LCS:", lcs(X, Y))
```

## **Explanation:**
- **Longest Common Subsequence:** Uses dynamic programming to find the longest subsequence present in both given sequences.

## **Python Example (Matrix Chain Multiplication):**

```python
def matrix_chain_multiplication(p):
    n = len(p) - 1
    dp = [[0] * n for _ in range(n)]

    for length in range(2, n + 1):
        for i in range(n - length + 1):
            j = i + length - 1
            dp[i][j] = float('inf')
            for k in range(i, j):
                q = dp[i][k] + dp[k + 1][j] + p[i] * p[k + 1] * p[j + 1]
                if q < dp[i][j]:
                    dp[i][j] = q

    return dp[0][n - 1]

# Example usage
p = [10, 20, 30, 40, 30]
print("Minimum number of multiplications:", matrix_chain_multiplication(p))
```

## **Explanation:**
- **Matrix Chain Multiplication:** Uses dynamic programming to find the most efficient way to multiply a given sequence of matrices.

