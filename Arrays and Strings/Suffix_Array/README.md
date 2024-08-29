# **Suffix Array**

## **Overview**

A Suffix Array is a data structure used to represent all suffixes of a string in a sorted order. It is useful for various string processing tasks such as substring search and pattern matching.

## **Where is it Used?**

- **String Matching:** Efficiently finding substrings and patterns.
- **Data Compression:** In algorithms like BWT (Burrows-Wheeler Transform).
- **Genome Analysis:** In bioinformatics for DNA sequence analysis.

## **How Does it Work?**

The suffix array is built by generating all suffixes of a string, sorting them, and storing their starting positions. This allows for efficient substring searching and other operations.

## **Python Example**

Here’s a basic example of constructing a suffix array for a string:

```python
def suffix_array_construction(s):
    suffixes = sorted((s[i:], i) for i in range(len(s)))
    return [suffix[1] for suffix in suffixes]

# Example usage
s = "banana"
suffix_array = suffix_array_construction(s)
print("Suffix Array:", suffix_array)  # Output: [5, 3, 1, 0, 4, 2]
```

- **Explanation:**  
  The function `suffix_array_construction` generates and sorts suffixes of the string `s`. It returns the starting indices of the suffixes in sorted order.

