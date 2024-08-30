# **Edit Distance**

## **Overview:**

The Edit Distance (Levenshtein Distance) is a measure of the difference between two sequences, computed as the minimum number of operations required to transform one sequence into the other.

## **Python Example:**

```python
def edit_distance(word1, word2):
    m, n = len(word1), len(word2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]

    for i in range(m + 1):
        for j in range(n + 1):
            if i == 0:
                dp[i][j] = j
            elif j == 0:
                dp[i][j] = i
            elif word1[i - 1] == word2[j - 1]:
                dp[i][j] = dp[i - 1][j - 1]
            else:
                dp[i][j] = 1 + min(dp[i - 1][j],    # Deletion
                                   dp[i][j - 1],    # Insertion
                                   dp[i - 1][j - 1]) # Substitution

    return dp[m][n]

# Example usage
word1 = "intention"
word2 = "execution"
print("Edit Distance:", edit_distance(word1, word2))
```

## **Explanation:**
- **Edit Distance:** Calculates the minimum number of edits required to convert one string into another.

