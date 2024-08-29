# **Selection Sort**

## **Overview**

Selection Sort is an in-place comparison-based algorithm. It works by dividing the list into a sorted and an unsorted part. It repeatedly selects the smallest (or largest) element from the unsorted part and moves it to the end of the sorted part.

## **Where is it Used?**

- **Small datasets:** Similar to Bubble Sort, it’s used for small datasets.
- **Memory-limited environments:** Since it requires no additional memory space, it’s useful in memory-limited environments.

## **How Does it Work?**

The algorithm repeatedly selects the minimum element from the unsorted portion and swaps it with the first unsorted element, expanding the sorted portion of the list by one element with each pass.

## **Python Example**

```python
def selection_sort(arr):
    n = len(arr)
    for i in range(n):
        min_idx = i
        for j in range(i + 1, n):
            if arr[j] < arr[min_idx]:
                min_idx = j
        # Swap the found minimum element with the first element
        arr[i], arr[min_idx] = arr[min_idx], arr[i]
    return arr

# Test the function
arr = [29, 10, 14, 37, 13]
print(selection_sort(arr))  # Output: [10, 13, 14, 29, 37]
```

- **Explanation:**  
  The `selection_sort` function finds the minimum element in the unsorted portion of the array and swaps it with the first element of the unsorted portion, effectively growing the sorted portion of the array by one element each time.
