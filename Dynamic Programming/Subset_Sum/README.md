# **Subset Sum**

## **Overview:**

The Subset Sum problem involves finding a subset of a given set that sums up to a specific target value.

## **Python Example:**

```python
def subset_sum(arr, target):
    n = len(arr)
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

    return dp[n][target]

# Example usage
arr = [3, 34, 4, 12, 5, 2]
target = 9
print("Subset with given sum exists:", subset_sum(arr, target))
```

## **Explanation:**
- **Subset Sum:** Determines if there is a subset of the array that sums up to the target value.

