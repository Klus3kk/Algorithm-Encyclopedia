# **Insertion Sort**

## **Overview**

Insertion Sort is a simple and intuitive comparison-based algorithm that builds the sorted array one element at a time. It works well for small or partially sorted datasets.

## **Where is it Used?**

- **Small datasets:** It’s efficient for small datasets.
- **Nearly sorted data:** It performs well when the array is already nearly sorted.

## **How Does it Work?**

Insertion Sort iterates over the list and, for each element, inserts it into its correct position in the already sorted portion of the list. This is done by shifting larger elements to the right to make space for the insertion.

## **Python Example**

```python
def insertion_sort(arr):
    for i in range(1, len(arr)):
        key = arr[i]
        j = i - 1
        # Move elements of arr[0...i-1], that are greater than key, to one position ahead
        while j >= 0 and key < arr[j]:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = key
    return arr

# Test the function
arr = [12, 11, 13, 5, 6]
print(insertion_sort(arr))  # Output: [5, 6, 11, 12, 13]
```

- **Explanation:**  
  The `insertion_sort` function takes each element and inserts it into its correct position in the sorted portion of the array by shifting elements as necessary.
