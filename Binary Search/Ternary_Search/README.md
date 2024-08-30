# **Ternary Search**

## **Overview:**

Ternary Search is a divide-and-conquer algorithm similar to binary search but divides the search space into three parts instead of two. It is used to find the maximum or minimum of a unimodal function.

## **Python Example:**

```python
def ternary_search(left, right, func):
    while right - left > 1e-9:  # Precision threshold
        mid1 = left + (right - left) / 3
        mid2 = right - (right - left) / 3
        
        if func(mid1) < func(mid2):
            left = mid1
        else:
            right = mid2
            
    return (left + right) / 2

# Example usage
def func(x):
    return -1 * (x - 2) ** 2 + 4  # Example function with a maximum

max_point = ternary_search(0, 4, func)
print("Maximum point:", max_point)
```

## **Explanation:**
- **Divide Space:** Divide the range into three parts.
- **Evaluate Function:** Compare the function values at the two division points.
- **Adjust Range:** Narrow down the search range based on the comparisons.

