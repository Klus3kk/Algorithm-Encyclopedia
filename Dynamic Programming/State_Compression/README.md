# **State Compression**

## **Overview:**

State Compression is used to optimize DP problems where the state space is too large. By compressing the state representation, the problem becomes more manageable.

## **Python Example:**

- **Subset Sum Problem with Bitmasking:**

```python
def subset_sum_bitmask(arr, target):
    n = len(arr)
    dp = [False] * (target + 1)
    dp[0] = True

    for num in arr:
        for i in range(target, num - 1, -1):
            dp[i] = dp[i] or dp[i - num]

    return dp[target]

# Example usage
arr = [3, 34, 4, 12, 5, 2]
target = 9
print("Subset sum possible:", subset_sum_bitmask(arr, target))
```

## **Explanation:**
- **Subset Sum with Bitmasking:** Uses a bitmask to compress the state space for the subset sum problem.

