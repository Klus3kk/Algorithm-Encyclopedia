# **Big Θ Notation**

## **Overview**

Big Θ notation provides a tight bound on the time or space complexity of an algorithm, meaning it describes both the upper and lower bounds. It gives a more precise characterization of an algorithm's complexity.

## **Where is it Used?**

- **Exact Analysis:** To understand the exact time or space complexity, considering both best and worst cases.
- **Balanced Comparison:** To compare algorithms that have similar upper and lower bounds.

## **How Does it Work?**

Big Θ notation is also expressed in terms of the input size `n`. For example:
- **Θ(1):** The algorithm has constant time complexity in both best and worst cases.
- **Θ(n):** The algorithm has linear time complexity in both best and worst cases.
- **Θ(n^2):** The algorithm has quadratic time complexity in both best and worst cases.

## **Python Example**

Consider the same example of summing elements in a list:

```python
def sum_of_elements(arr):
    total = 0
    for element in arr:
        total += element
    return total

arr = [1, 2, 3, 4, 5]
print(sum_of_elements(arr))
```

- **Complexity Analysis:**  
  The time complexity for this function is Θ(n) because the function will always require linear time, no matter the input.
