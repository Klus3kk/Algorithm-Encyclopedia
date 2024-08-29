# **Bubble Sort**

## **Overview**

Bubble Sort is a simple comparison-based sorting algorithm. It repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. This process is repeated until the list is sorted.

## **Where is it Used?**

- **Small datasets:** Due to its simplicity, Bubble Sort is suitable for small datasets.
- **Learning purposes:** It’s often used as an introductory example of sorting algorithms.

## **How Does it Work?**

Bubble Sort works by repeatedly comparing each pair of adjacent elements in the list and swapping them if they are out of order. The algorithm continues until no swaps are needed, indicating that the list is sorted.

## **Python Example**

```python
def bubble_sort(arr):
    n = len(arr)
    for i in range(n):
        swapped = False
        for j in range(0, n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                swapped = True
        # If no elements were swapped, the array is sorted
        if not swapped:
            break
    return arr

# Test the function
arr = [64, 34, 25, 12, 22, 11, 90]
print(bubble_sort(arr))  # Output: [11, 12, 22, 25, 34, 64, 90]
```

- **Explanation:**  
  The `bubble_sort` function iteratively compares and swaps adjacent elements if they are in the wrong order. The algorithm's name comes from the way larger elements "bubble" to the top of the list.
