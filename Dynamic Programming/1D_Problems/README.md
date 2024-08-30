# **1D Problems**

## **Overview:**

1D dynamic programming problems typically involve a single dimension of data, such as an array, and require optimization over a linear sequence.

## **Python Example:**

- **Fibonacci Sequence:**

```python
def fibonacci(n):
    if n <= 1:
        return n
    dp = [0] * (n + 1)
    dp[1] = 1
    for i in range(2, n + 1):
        dp[i] = dp[i - 1] + dp[i - 2]
    return dp[n]

# Example usage
n = 10
print("Fibonacci number at position", n, "is:", fibonacci(n))
```

- **Climbing Stairs:**

```python
def climb_stairs(n):
    if n <= 1:
        return 1
    dp = [0] * (n + 1)
    dp[0] = 1
    dp[1] = 1
    for i in range(2, n + 1):
        dp[i] = dp[i - 1] + dp[i - 2]
    return dp[n]

# Example usage
n = 5
print("Number of ways to climb", n, "stairs:", climb_stairs(n))
```

## **Explanation:**
- **Fibonacci Sequence:** Uses DP to compute the nth Fibonacci number efficiently by storing intermediate results.
- **Climbing Stairs:** Calculates the number of ways to reach the top by considering the number of ways to reach each step.

