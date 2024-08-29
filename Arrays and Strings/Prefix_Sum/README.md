# **Prefix Sum**

## **Overview**

The Prefix Sum Technique involves preprocessing an array to create a new array where each element is the sum of all previous elements in the original array. This allows for efficient querying of subarray sums.

## **Where is it Used?**

- **Range Sum Queries:** Quickly finding the sum of elements between two indices.
- **Subarray Problems:** Efficiently solving problems that require repeated sum queries.

## **How Does it Work?**

A prefix sum array is built where each element at index `i` contains the sum of the array elements from the start up to index `i`. This allows for quick calculation of subarray sums using the formula:

\[ \text{sum}(i, j) = \text{prefix}[j] - \text{prefix}[i-1] \]

## **Python Example**

Here’s an example of how to compute range sum queries using the Prefix Sum Technique:

```python
def build_prefix_sum(arr):
    prefix_sum = [0] * len(arr)
    prefix_sum[0] = arr[0]
    for i in range(1, len(arr)):
        prefix_sum[i] = prefix_sum[i - 1] + arr[i]
    return prefix_sum

def range_sum(prefix_sum, i, j):
    if i == 0:
        return prefix_sum[j]
    return prefix_sum[j] - prefix_sum[i - 1]

# Example usage
arr = [1, 2, 3, 4, 5]
prefix_sum = build_prefix_sum(arr)
print("Sum of elements from index 1 to 3:", range_sum(prefix_sum, 1, 3))  # Output: 9
```

- **Explanation:**  
  The `build_prefix_sum` function creates a prefix sum array. The `range_sum` function uses this array to quickly compute the sum of elements between two indices.

