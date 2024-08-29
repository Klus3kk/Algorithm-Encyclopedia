# **Sliding Window**

## **Overview**

The Sliding Window Technique is used to find a subset of elements or a contiguous segment in an array or list that satisfies certain conditions. It is particularly effective for problems involving subarrays of fixed or variable length.

## **Where is it Used?**

- **Maximum/Minimum Sum Subarray:** Finding the maximum or minimum sum of a subarray of fixed length.
- **Substring Problems:** Finding substrings with specific properties (e.g., containing all characters of another string).

## **How Does it Work?**

The technique involves maintaining a window (a subset of the array) that slides over the data structure. The window size can be fixed or variable, and it moves to cover different segments of the array.

## **Python Example**

Here’s an example of the Sliding Window Technique to find the maximum sum of a subarray of length `k`:

```python
def max_sum_subarray(arr, k):
    n = len(arr)
    if n < k:
        return None  # Not enough elements for the window size

    window_sum = sum(arr[:k])
    max_sum = window_sum

    for i in range(k, n):
        window_sum += arr[i] - arr[i - k]
        max_sum = max(max_sum, window_sum)

    return max_sum

# Example usage
arr = [1, 2, 3, 4, 5, 6, 7]
k = 3
print("Maximum sum of subarray of length", k, "is:", max_sum_subarray(arr, k))  # Output: 18
```

- **Explanation:**  
  The algorithm initializes the sum of the first window and then slides the window one element at a time. It updates the sum by adding the new element and subtracting the element that is no longer in the window, keeping track of the maximum sum.
