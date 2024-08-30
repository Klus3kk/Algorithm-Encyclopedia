# **Matrix Chain Multiplication**

## **Overview:**

Matrix Chain Multiplication involves finding the most efficient way to multiply a given sequence of matrices to minimize the number of scalar multiplications.

## **Python Example:**

```python
def matrix_chain_order(p):
    n = len(p) - 1
    dp = [[float('inf')] * n for _ in range(n)]
    for i in range(n):
        dp[i][i] = 0

    for length in range(2, n + 1):
        for i in range(n - length + 1):
            j = i + length - 1
            for k in range(i, j):
                dp[i][j] = min(dp[i][j], dp[i][k] + dp[k + 1][j] + p[i] * p[k + 1] * p[j + 1])

    return dp[0][n - 1]

# Example usage
p = [10, 20, 30, 40, 30]
print("Minimum number of multiplications is:", matrix_chain_order(p))
```

## **Explanation:**
- **Matrix Chain Multiplication:** Determines the optimal way to parenthesize a sequence of matrices to minimize the number of multiplications.

