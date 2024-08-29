# **Two Pointers Technique**

## **Overview**

The Two Pointers Technique is a strategy for solving problems by using two pointers or indices to traverse the data structure, often arrays or lists. It is particularly useful for problems involving pairs or subarrays.

## **Where is it Used?**

- **Pair Problems:** Finding pairs with a specific sum in a sorted array.
- **Subarray Problems:** Finding subarrays meeting certain criteria (e.g., minimum length with a sum greater than a given value).
- **String Problems:** Checking for palindrome properties or substrings.

## **How Does it Work?**

Two pointers are initialized at different positions in the array and move towards each other or in a specific direction based on the problem's requirements.

## **Python Example**

Here’s an example of the Two Pointers Technique to find if a pair with a given sum exists in a sorted array:

```python
def has_pair_with_sum(arr, target_sum):
    left, right = 0, len(arr) - 1

    while left < right:
        current_sum = arr[left] + arr[right]
        if current_sum == target_sum:
            return True
        elif current_sum < target_sum:
            left += 1
        else:
            right -= 1

    return False

# Example usage
arr = [1, 2, 3, 4, 5, 6]
target_sum = 8
print("Pair with sum exists:", has_pair_with_sum(arr, target_sum))  # Output: True
```

- **Explanation:**  
  The algorithm initializes two pointers: `left` at the beginning and `right` at the end of the array. By adjusting the pointers based on the sum of the elements they point to, it efficiently finds pairs that meet the target sum.
