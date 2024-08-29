# **Radix Sort**

## **Overview**

Radix Sort is a non-comparison-based sorting algorithm that sorts numbers digit by digit, starting from the least significant digit to the most significant digit, using a stable subroutine like Counting Sort.

## **Where is it Used?**

- **Large integers:** Suitable for sorting large integers.
- **Linear time sorting:** It works in O(nk) time, where `n` is the number of elements and `k` is the number of digits.

## **How Does it Work?**

Radix Sort processes each digit of the numbers in the list, starting from the least significant digit, and uses a stable sorting algorithm like Counting Sort to sort based on that digit.

## **Python Example**

```python
def counting_sort_for_radix(arr, exp):
    n = len(arr)
    output = [0] * n
    count = [0] * 10

    for i in range(n):
        index = arr[i] // exp
        count[index % 10] += 1

    for i in range(1, 10):
        count[i] += count[i - 1]

    i = n - 1
    while i >= 0:
        index = arr[i] // exp
        output[count[index % 10] - 1] = arr[i]
        count[index % 10] -= 1
        i -= 1

    for i in range(n):
        arr[i] = output[i]

def radix_sort(arr):
    max_val = max(arr)
    exp = 1
    while max_val // exp > 0:
        counting_sort_for_radix(arr, exp)
        exp *= 10
    return arr

# Test the function
arr = [170, 45, 75, 90, 802, 24, 2, 66]
print(radix_sort(arr))  # Output: [2, 24, 45, 66, 75, 90, 170, 802]
```

- **Explanation:**  
  The `radix_sort` function sorts the array by processing each digit, starting from the least significant digit, using Counting Sort as a subroutine.
