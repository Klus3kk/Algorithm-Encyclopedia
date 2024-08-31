# **Z-Algorithm**

## **Overview:**

The Z-Algorithm computes the Z-array of a string, where each position in the Z-array indicates the length of the longest substring starting from that position which is also a prefix of the string. This is useful for string matching and pattern search.

## **Python Example:**

```python
def z_algorithm(s):
    z = [0] * len(s)
    l, r, k = 0, 0, 0
    for i in range(1, len(s)):
        if i > r:
            l, r = i, i
            while r < len(s) and s[r] == s[r - l]:
                r += 1
            z[i] = r - l
            r -= 1
        else:
            k = i - l
            if z[k] < r - i + 1:
                z[i] = z[k]
            else:
                l = i
                while r < len(s) and s[r] == s[r - l]:
                    r += 1
                z[i] = r - l
                r -= 1
    return z

def z_search(text, pattern):
    concatenated = pattern + '$' + text
    z = z_algorithm(concatenated)
    pattern_length = len(pattern)
    for i in range(pattern_length + 1, len(z)):
        if z[i] == pattern_length:
            return i - pattern_length - 1
    return -1

# Example usage
text = "ababcababcabc"
pattern = "abc"
print("Pattern found at index:", z_search(text, pattern))
```

## **Explanation:**
- **Z-Algorithm:** Efficiently computes Z-values for a string, aiding in pattern search by concatenating the pattern with the text.

