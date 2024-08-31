# **Suffix Trees**

## **Overview:**

Suffix trees are advanced data structures used for solving string problems. They allow for efficient substring searches, longest repeated substring discovery, and longest common substring calculations by constructing a tree of all suffixes of a string.

## **Python Example:**

```python
class SuffixTree:
    def __init__(self, s):
        self.s = s
        self.build_suffix_tree()
    
    def build_suffix_tree(self):
        self.suffixes = [self.s[i:] for i in range(len(self.s))]
        self.suffixes.sort()
    
    def search(self, pattern):
        left, right = 0, len(self.suffixes) - 1
        while left <= right:
            mid = (left + right) // 2
            if self.suffixes[mid].startswith(pattern):
                return mid
            elif self.suffixes[mid] < pattern:
                left = mid + 1
            else:
                right = mid - 1
        return -1

# Example usage
st = SuffixTree("banana")
print("Suffixes:", st.suffixes)
print("Pattern 'ana' found at index:", st.search("ana"))
```

## **Explanation:**
- **Suffix Trees:** Construct and sort suffixes of a string for efficient pattern searching and substring analysis.

