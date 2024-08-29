# **Heap Sort**

## **Overview**

Heap Sort is a comparison-based algorithm that uses a binary heap data structure. It’s similar to selection sort but uses a heap to find the maximum or minimum element efficiently.

## **Where is it Used?**

- **Priority queues:** Often used in implementing priority queues.
- **In-place sorting:** It doesn’t require additional

 storage space.

## **How Does it Work?**

Heap Sort first builds a max-heap from the input data, then repeatedly extracts the maximum element from the heap and rebuilds the heap until all elements are sorted.

## **Python Example**

```python
def heapify(arr, n, i):
    largest = i
    left = 2 * i + 1
    right = 2 * i + 2

    if left < n and arr[i] < arr[left]:
        largest = left

    if right < n and arr[largest] < arr[right]:
        largest = right

    if largest != i:
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)

def heap_sort(arr):
    n = len(arr)

    for i in range(n // 2 - 1, -1, -1):
        heapify(arr, n, i)

    for i in range(n - 1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]
        heapify(arr, i, 0)

    return arr

# Test the function
arr = [12, 11, 13, 5, 6, 7]
print(heap_sort(arr))  # Output: [5, 6, 7, 11, 12, 13]
```

- **Explanation:**  
  The `heap_sort` function first builds a max-heap, then repeatedly extracts the largest element from the heap, swaps it with the last unsorted element, and re-heapifies the array.
