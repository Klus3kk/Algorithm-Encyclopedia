# **Heap Sort**

## **Overview:**

Heap Sort is a comparison-based sorting algorithm that leverages the heap data structure to sort elements. It involves building a heap from the data and then repeatedly extracting the maximum (or minimum) element from the heap and reconstructing the heap.

## **Python Example:**

```python
def heap_sort(arr):
    # Build a max-heap
    n = len(arr)
    for i in range(n // 2 - 1, -1, -1):
        heapify(arr, n, i)

    # Extract elements one by one
    for i in range(n - 1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]  # Swap
        heapify(arr, i, 0)  # Heapify the root of the tree

def heapify(arr, n, i):
    largest = i
    left = 2 * i + 1
    right = 2 * i + 2

    if left < n and arr[i] < arr[left]:
        largest = left

    if right < n and arr[largest] < arr[right]:
        largest = right

    if largest != i:
        arr[i], arr[largest] = arr[largest], arr[i]  # Swap
        heapify(arr, n, largest)

# Example usage
arr = [12, 11, 13, 5, 6, 7]
heap_sort(arr)
print("Sorted array:", arr)
```

## **Explanation:**
- **Build Max-Heap:** Convert the array into a max-heap.
- **Extract Elements:** Swap the root with the last element, reduce heap size, and restore heap order.
