# **Top K Elements Problems**

## **Overview:**

Finding the top K elements involves retrieving the K largest or smallest elements from a collection. Heaps are commonly used to solve this problem efficiently.

## **Python Example:**

```python
import heapq

def top_k_elements(arr, k):
    return heapq.nlargest(k, arr)

# Example usage
arr = [3, 2, 1, 5, 6, 4]
print("Top 3 elements:", top_k_elements(arr, 3))
```

## **Explanation:**
- **`heapq.nlargest`:** Retrieves the K largest elements from the array.

