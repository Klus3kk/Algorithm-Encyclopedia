# **Knapsack Problem**

## **Overview:**

The Knapsack Problem involves selecting items with given weights and values to maximize the total value in a knapsack of fixed capacity. The problem comes in two main forms: 0/1 Knapsack and Fractional Knapsack.

## **Python Example (0/1 Knapsack):**

```python
def knapsack(weights, values, capacity):
    n = len(values)
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
- **0/1 Knapsack:** Uses dynamic programming to build up the maximum value for each weight up to the capacity, considering each item only once.

