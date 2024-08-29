# **Subarray Problems**

## **Overview**

Subarray problems involve finding specific subarrays within an array that meet certain criteria. These problems can be addressed using various techniques depending on the constraints and requirements.

## **Where is it Used?**

- **Maximum/Minimum Subarray Sum:** Finding subarrays with the highest or lowest sum.
- **Fixed-Length Subarrays:** Working with subarrays of a specific length.
- **Contiguous Subarrays:** Handling problems involving contiguous sequences of elements.

## **How Does it Work?**

The approach depends on the specific problem but often involves iterating over possible subarrays and applying optimization techniques to find the desired result.

## **Python Example**

Here’s an example of finding the maximum sum of a contiguous subarray using Kadane’s Algorithm:

```python
def max_sum_contiguous_subarray(arr):
    max_current = max_global = arr[0]

    for i in range(1, len(arr)):
        max_current = max(arr[i], max_current + arr[i])
        if max_current > max_global:
            max_global = max_current

    return max_global

# Example usage
arr = [1, -2, 3, 4, -1, 2, 1, -5, 4]
print("Maximum sum of contiguous subarray:", max_sum_contiguous_subarray(arr))  # Output: 9
```

- **Explanation:**  
  Kadane’s Algorithm efficiently computes the maximum sum of a contiguous subarray by maintaining the maximum sum ending at the current position and updating the global maximum.

