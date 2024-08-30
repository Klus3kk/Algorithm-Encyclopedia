# **Binary Search on Sorted Arrays**

## **Overview:**

Binary Search on Sorted Arrays is used to find a target value in a sorted array. The algorithm divides the array into two halves and eliminates one half from consideration, based on the comparison of the target value with the middle element.

## **Python Example:**

```python
def binary_search(arr, target):
    left, right = 0, len(arr) - 1
    
    while left <= right:
        mid = left + (right - left) // 2
        
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
            
    return -1

# Example usage
arr = [2, 3, 4, 10, 40]
target = 10
print("Index of target:", binary_search(arr, target))
```

## **Explanation:**
- **Initialization:** Start with `left` at the beginning and `right` at the end of the array.
- **Middle Element:** Calculate the middle index and compare the target with the middle element.
- **Adjust Range:** Narrow down the search range based on the comparison.

