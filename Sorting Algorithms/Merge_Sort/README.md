# **Merge Sort**

## **Overview**

Merge Sort is a divide-and-conquer algorithm that splits the list into smaller sublists, sorts each sublist, and then merges the sorted sublists back together. It has a consistent time complexity of O(n log n).

## **Where is it Used?**

- **Large datasets:** It’s efficient for large datasets.
- **Stable sorting:** It preserves the relative order of equal elements.

## **How Does it Work?**

Merge Sort recursively divides the array into two halves until each sublist contains a single element. Then, it merges the sublists back together in a sorted order.

## **Python Example**

```python
def merge_sort(arr):
    if len(arr) > 1:
        mid = len(arr) // 2
        left_half = arr[:mid]
        right_half = arr[mid:]

        merge_sort(left_half)
        merge_sort(right_half)

        i = j = k = 0

        # Merge the two halves
        while i < len(left_half) and j < len(right_half):
            if left_half[i] < right_half[j]:
                arr[k] = left_half[i]
                i += 1
            else:
                arr[k] = right_half[j]
                j += 1
            k += 1

        # Check if any element was left
        while i < len(left_half):
            arr[k] = left_half[i]
            i += 1
            k += 1

        while j < len(right_half):
            arr[k] = right_half[j]
            j += 1
            k += 1

    return arr

# Test the function
arr = [38, 27, 43, 3, 9, 82, 10]
print(merge_sort(arr))  # Output: [3, 9, 10, 27, 38, 43, 82]
```

- **Explanation:**  
  The `merge_sort` function recursively splits the array into halves, sorts each half, and then merges them back together in the correct order.

