# **Quick Sort**

## **Overview**

Quick Sort is a highly efficient divide-and-conquer algorithm that works by selecting a "pivot" element from the array and partitioning the other elements into two sub-arrays, according to whether they are less than or greater than the pivot.

## **Where is it Used?**

- **Large datasets:** It’s efficient for large datasets.
- **In-place sorting:** It doesn’t require additional storage space.

## **How Does it Work?**

Quick Sort partitions the array around a pivot element, such that elements less than the pivot are on one side and elements greater than the pivot are on the other. It then recursively sorts the sub-arrays.

## **Python Example**

```python
def quick_sort(arr):
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return quick_sort(left) + middle + quick_sort(right)

# Test the function
arr = [3, 6, 8, 10, 1, 2, 1]
print(quick_sort(arr))  # Output: [1, 1, 2, 3, 6, 8, 10]
```

- **Explanation:**  
  The `quick_sort` function selects a pivot element, then partitions the array into elements less than, equal to, and greater than the pivot, and finally recursively sorts the partitions.
