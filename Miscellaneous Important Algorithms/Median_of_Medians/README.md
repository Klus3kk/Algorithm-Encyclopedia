# **Median of Medians**

## **Overview:**

Median of Medians is a deterministic algorithm to find the k-th smallest element in linear time. It improves upon Quickselect by ensuring better worst-case performance.

## **Python Example:**

```python
def median_of_medians(arr, k):
    if len(arr) <= 5:
        return sorted(arr)[k]

    medians = [sorted(arr[i:i+5])[len(arr[i:i+5])//2] for i in range(0, len(arr), 5)]
    pivot = median_of_medians(medians, len(medians)//2)
    
    lows = [el for el in arr if el < pivot]
    highs = [el for el in arr if el > pivot]
    pivots = [el for el in arr if el == pivot]
    
    if k < len(lows):
        return median_of_medians(lows, k)
    elif k < len(lows) + len(pivots):
        return pivots[0]
    else:
        return median_of_medians(highs, k - len(lows) - len(pivots))

# Example usage
arr = [3, 6, 8, 10, 1, 2, 1]
k = 3
print(f"{k}-th smallest element is:", median_of_medians(arr, k-1))
```

## **Explanation:**
- **Median of Medians:** Partitions the array based on the median of medians, ensuring better performance guarantees compared to Quickselect.
