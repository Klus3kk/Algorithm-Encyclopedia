# **2D Problems**

## **Overview:**

2D dynamic programming problems involve two dimensions of data, often represented as a grid or table. These problems require optimization over a 2D structure.

## **Python Example:**

- **Knapsack Problem:**

```python
def knapsack(weights, values, W):
    n = len(weights)
    dp = [[0] * (W + 1) for _ in range(n + 1)]

    for i in range(1, n + 1):
        for w in range(W + 1):
            if weights[i - 1] <= w:
                dp[i][w] = max(dp[i - 1][w], dp[i - 1][w - weights[i - 1]] + values[i - 1])
            else:
                dp[i][w] = dp[i - 1][w]

    return dp[n][W]

# Example usage
weights = [1, 2, 3]
values = [60, 100, 120]
W = 5
print("Maximum value in Knapsack:", knapsack(weights, values, W))
```

- **Coin Change:**

```python
def coin_change(coins, amount):
    dp = [float('inf')] * (amount + 1)
    dp[0] = 0
    for i in range(1, amount + 1):
        for coin in coins:
            if i - coin >= 0:
                dp[i] = min(dp[i], dp[i - coin] + 1)
    return dp[amount] if dp[amount] != float('inf') else -1

# Example usage
coins = [1, 2, 5]
amount = 11
print("Minimum number of coins needed:", coin_change(coins, amount))
```

## **Explanation:**
- **Knapsack Problem:** Solves the problem of selecting items to maximize value without exceeding capacity.
- **Coin Change:** Determines the minimum number of coins needed to make a given amount.
