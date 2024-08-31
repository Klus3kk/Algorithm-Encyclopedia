# **Quickselect**

## **Overview:**

Quickselect is an algorithm to find the k-th smallest (or largest) element in an unordered list. It is related to Quick Sort but focuses only on finding a specific element.

## **Python Example:**

```python
import random

def quickselect(arr, k):
    if len(arr) == 1:
        return arr[0]
    
    pivot = random.choice(arr)
    lows = [el for el in arr if el < pivot]
    highs = [el for el in arr if el > pivot]
    pivots = [el for el in arr if el == pivot]
    
    if k < len(lows):
        return quickselect(lows, k)
    elif k < len(lows) + len(pivots):
        return pivots[0]
    else:
        return quickselect(highs, k - len(lows) - len(pivots))

# Example usage
arr = [3, 6, 8, 10, 1, 2, 1]
k = 3
print(f"{k}-th smallest element is:", quickselect(arr, k-1))
```

## **Explanation:**
- **Quickselect:** Chooses a pivot and partitions the array into elements less than, equal to, and greater than the pivot. It then recurses into the relevant partition to find the k-th smallest element.

