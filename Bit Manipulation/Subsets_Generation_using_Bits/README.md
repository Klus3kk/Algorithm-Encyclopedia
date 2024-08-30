# **Subsets Generation using Bits**

## **Overview:**

Subsets generation using bits involves using bitwise operations to generate all possible subsets of a given set.

## **Python Example:**

- **Generate All Subsets:**

```python
def generate_subsets(arr):
    subsets = []
    n = len(arr)
    for i in range(1 << n):  # 2^n subsets
        subset = []
        for j in range(n):
            if i & (1 << j):
                subset.append(arr[j])
        subsets.append(subset)
    return subsets

# Example usage
arr = [1, 2, 3]
print("All subsets:", generate_subsets(arr))
```

## **Explanation:**
- **Generate Subsets:** Uses bitwise operations to iterate through all possible combinations of elements and generates subsets.

