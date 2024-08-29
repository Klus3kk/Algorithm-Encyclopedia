# **Big Ω Notation**

## **Overview**

Big Ω notation provides a lower bound on the time or space complexity of an algorithm. It describes the best-case scenario, helping to understand the minimum time or space an algorithm will require.

## **Where is it Used?**

- **Best-Case Analysis:** To determine the best-case time or space complexity.
- **Algorithm Guarantees:** To understand the guaranteed minimum performance.

## **How Does it Work?**

Big Ω notation is also expressed in terms of the input size `n`. For example:
- **Ω(1):** Constant time in the best case.
- **Ω(n):** Linear time in the best case.
- **Ω(n^2):** Quadratic time in the best case.

## **Python Example**

For the same function `sum_of_elements`, in the best-case scenario where all elements are summed, the operation still takes linear time:

```python
def sum_of_elements(arr):
    return sum(arr)  # Python's built-in sum function also runs in O(n) time.

arr = [1, 2, 3, 4, 5]
print(sum_of_elements(arr))
```

- **Complexity Analysis:**  
  Even in the best case, the function requires Ω(n) time, as each element must be visited at least once.
