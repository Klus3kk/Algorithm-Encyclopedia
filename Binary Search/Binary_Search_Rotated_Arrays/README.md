# **Binary Search on Rotated Arrays**

## **Overview:**

Binary Search on Rotated Arrays handles arrays that have been rotated at some pivot. The algorithm adjusts the search strategy to account for the rotation while still maintaining efficiency.

## **Python Example:**

```python
def binary_search_rotated(arr, target):
    left, right = 0, len(arr) - 1
    
    while left <= right:
        mid = left + (right - left) // 2
        
        if arr[mid] == target:
            return mid
        
        # Determine which part is sorted
        if arr[left] <= arr[mid]:
            if arr[left] <= target < arr[mid]:
                right = mid - 1
            else:
                left = mid + 1
        else:
            if arr[mid] < target <= arr[right]:
                left = mid + 1
            else:
                right = mid - 1
                
    return -1

# Example usage
arr = [15, 18, 2, 3, 6, 12]
target = 3
print("Index of target:", binary_search_rotated(arr, target))
```

## **Explanation:**
- **Determine Sorted Part:** Identify which half of the array is sorted.
- **Adjust Search Range:** Decide which half to search based on the target and sorted part.

