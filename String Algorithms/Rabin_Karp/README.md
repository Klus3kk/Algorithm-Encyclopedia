# **Rabin-Karp**

## **Overview:**

The Rabin-Karp algorithm is a string search algorithm that uses hashing to find any one of a set of pattern strings in a text. It computes hash values for the pattern and the substrings of the text to identify matches efficiently. It's particularly effective for searching multiple patterns simultaneously.

## **Python Example:**

```python
def rabin_karp(text, pattern, d=256, q=101):
    def hash_function(s, end):
        h = 0
        for i in range(end):
            h = (d * h + ord(s[i])) % q
        return h
    
    def roll_hash(h, old, new, m, d, q):
        h = (d * (h - ord(old) * h_m) + ord(new)) % q
        return h
    
    n, m = len(text), len(pattern)
    if m > n:
        return -1
    
    h_pattern = hash_function(pattern, m)
    h_text = hash_function(text, m)
    h_m = pow(d, m - 1, q)
    
    for i in range(n - m + 1):
        if h_text == h_pattern and text[i:i + m] == pattern:
            return i
        
        if i < n - m:
            h_text = roll_hash(h_text, text[i], text[i + m], m, d, q)
    
    return -1

# Example usage
text = "ababcababcabc"
pattern = "abc"
print("Pattern found at index:", rabin_karp(text, pattern))
```

## **Explanation:**
- **Rabin-Karp Algorithm:** Utilizes hashing for efficient pattern searching, particularly useful when dealing with multiple patterns.

