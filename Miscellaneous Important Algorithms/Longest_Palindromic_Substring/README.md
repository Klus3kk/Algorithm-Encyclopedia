# **Manacher's Algorithm (Longest Palindromic Substring)**

## **Overview:**

Manacher's Algorithm finds the longest palindromic substring in linear time. It is an efficient algorithm that expands around possible centers and utilizes previously computed results to optimize performance.

## **Python Example:**

```python
def manacher(s):
    s = '#' + '#'.join(s) + '#'
    n = len(s)
    lps = [0] * n
    center = right = 0
    for i in range(n):
        mirror = 2 * center - i
        if i < right:
            lps[i] = min(right - i, lps[mirror])
        while i + lps[i] + 1 < n and i - lps[i] - 1 >= 0 and s[i + lps[i] + 1] == s[i - lps[i] - 1]:
            lps[i] += 1
        if i + lps[i] > right:
            center = i
            right = i + lps[i]
    
    max_len, max_center = max((n, i) for i, n in enumerate(lps))
    return s[max_center - max_len + 1 : max_center + max_len].replace('#', '')

# Example usage
print("Longest Palindromic Substring:", manacher("babad"))
```

## **Explanation:**
- **Manacher's Algorithm:** Transforms the string to handle even-length palindromes, computes the longest palindromic substring using centers and expands while maintaining the longest found so far.

