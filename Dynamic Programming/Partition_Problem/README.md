# **Partition Problem**

## **Overview:**

The Partition Problem involves dividing a set into two subsets such that the difference between their sums is minimized.

## **Python Example:**

```python
def partition_problem(arr):
    total_sum = sum(arr)
    n = len(arr)
    target = total_sum // 2
    dp = [[False] * (target + 1) for _ in range(n + 1)]
    dp[0][0] = True

    for i in range(1, n + 1):
        dp[i][0] = True

    for i in range(1, n + 1):
        for j in range(1, target + 1):
            if arr[i - 1] <= j:
                dp[i][j] = dp[i - 1][j] or dp[i - 1][j - arr[i - 1]]
            else:
                dp[i][j] = dp[i - 1][j]

    return total_sum - 2 * max(j for j in range(target + 1) if dp[n][j])

# Example usage
arr = [1, 5, 11, 5]
print("Minimum subset sum difference:", partition_problem(arr))
```

## **Explanation:**
- **Partition Problem:** Finds the minimum difference between the sums of two subsets by partitioning the set.

