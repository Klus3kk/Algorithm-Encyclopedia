# **KMP (Knuth-Morris-Pratt)**

## **Overview:**

The KMP algorithm is designed to efficiently find the occurrence of a pattern within a text. It preprocesses the pattern to create a partial match table (or "prefix function") that allows the search to skip over parts of the text where the pattern will not match, resulting in a linear time complexity.

## **Python Example:**

```python
def kmp_search(text, pattern):
    def build_lps_array(pattern):
        lps = [0] * len(pattern)
        length = 0
        i = 1
        while i < len(pattern):
            if pattern[i] == pattern[length]:
                length += 1
                lps[i] = length
                i += 1
            else:
                if length != 0:
                    length = lps[length - 1]
                else:
                    lps[i] = 0
                    i += 1
        return lps
    
    lps = build_lps_array(pattern)
    i = j = 0
    while i < len(text):
        if pattern[j] == text[i]:
            i += 1
            j += 1
        
        if j == len(pattern):
            return i - j  # Found pattern at index i - j
        elif i < len(text) and pattern[j] != text[i]:
            if j != 0:
                j = lps[j - 1]
            else:
                i += 1
    return -1  # Pattern not found

# Example usage
text = "ababcababcabc"
pattern = "abc"
print("Pattern found at index:", kmp_search(text, pattern))
```

## **Explanation:**
- **KMP Algorithm:** Efficiently searches for a pattern within a text by preprocessing the pattern to skip unnecessary comparisons.

