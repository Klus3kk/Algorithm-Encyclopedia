# **Counting Sort**

## **Overview**

Counting Sort is a non-comparison-based algorithm that sorts integers by counting the occurrences of each unique value in the array. It works well for integers within a limited range.

## **Where is it Used?**

- **Fixed range integers:** Useful when sorting integers within a specific range.
- **Linear time sorting:** It runs in O(n + k) time, where `k` is the range of the input values.

## **How Does it Work?**

Counting Sort creates a count array to store the frequency of each unique value in the input array, then calculates the positions of each value in the sorted array.

## **Python Example**

```python
def counting_sort(arr):
    max_val = max(arr)
    count = [0] * (max_val + 1)
    output = [0] * len(arr)

    for num in arr:
        count[num] += 1

    for i in range(1, len(count)):
        count[i] += count[i - 1]

    for num in reversed(arr):
        output[count[num] - 1] = num
        count[num] -= 1

    return output

# Test the function
arr = [4, 2, 2, 8, 3, 3, 1]
print(counting_sort(arr))  # Output: [1, 2, 2, 3, 3, 4, 8]
```

- **Explanation:**  
  The `counting_sort` function creates a count array to track the frequency of each value, then uses this array to place each element in its correct sorted position.
