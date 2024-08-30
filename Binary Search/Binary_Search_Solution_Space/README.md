# **Binary Search in Solution Space**

## **Overview:**

Binary Search in Solution Space is used for problems where the search space is continuous, and you need to find an optimal solution by applying binary search principles to the range of possible solutions.

## **Python Example:**

```python
def find_min_sqrt(n):
    left, right = 0, n
    
    while left <= right:
        mid = left + (right - left) // 2
        square = mid * mid
        
        if square == n:
            return mid
        elif square < n:
            left = mid + 1
        else:
            right = mid - 1
            
    return right

# Example usage
n = 16
print("Integer square root:", find_min_sqrt(n))
```

## **Explanation:**
- **Initialization:** Set `left` to 0 and `right` to the number `n`.
- **Middle Calculation:** Compute the square of the middle value.
- **Adjust Range:** Adjust the search range based on the comparison with `n`.

