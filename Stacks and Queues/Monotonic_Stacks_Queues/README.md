# **Monotonic Stacks and Queues**

## **Overview**

Monotonic stacks and queues are specialized data structures where elements are maintained in a non-decreasing or non-increasing order. They are used to solve problems that involve maintaining order constraints.

## **Where is it Used?**

- **Next Greater Element:** Finding the next greater element for each element in an array.
- **Sliding Window Maximum:** Tracking maximum values in a sliding window.

## **How Does it Work?**

- **Monotonic Stack:** Maintains a stack where elements are in a non-decreasing or non-increasing order.
- **Monotonic Queue:** Maintains a queue with similar constraints.

## **Python Example**

Here’s an example of using a monotonic stack to find the next greater element:

```python
def next_greater_elements(nums):
    stack = []
    result = [-1] * len(nums)
    for i in range(len(nums)):
        while stack and nums[stack[-1]] < nums[i]:
            result[stack.pop()] = nums[i]
        stack.append(i)
    return result

# Example usage
nums = [4, 5, 2, 10, 8]
print("Next Greater Elements:", next_greater_elements(nums))  # Output: [5, 10, 10, -1, -1]
```

- **Explanation:**  
  The `next_greater_elements` function finds the next greater element for each element in the list using a monotonic stack.

