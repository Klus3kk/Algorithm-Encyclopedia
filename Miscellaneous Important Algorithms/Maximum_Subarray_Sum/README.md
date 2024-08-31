# **Kadane’s Algorithm (Maximum Subarray Sum)**

## **Overview:**

Kadane’s Algorithm finds the maximum sum of a contiguous subarray within a one-dimensional array of numbers. It is a dynamic programming approach that runs in linear time.

## **Python Example:**

```python
def kadane(arr):
    max_ending_here = max_so_far = arr[0]
    for x in arr[1:]:
        max_ending_here = max(x, max_ending_here + x)
        max_so_far = max(max_so_far, max_ending_here)
    return max_so_far

# Example usage
print("Maximum subarray sum:", kadane([1, -2, 3, 4, -1, 2, 1, -5, 4]))
```

## **Explanation:**
- **Kadane’s Algorithm:** Iterates through the array while maintaining the maximum subarray sum ending at the current position and updating the global maximum.

