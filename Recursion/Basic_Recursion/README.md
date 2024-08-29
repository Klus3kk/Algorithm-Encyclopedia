# **Basic Recursion**

## **Overview**

Basic recursion involves a function calling itself with modified arguments until a base case is reached, which halts further recursive calls. Each recursive call works on a smaller instance of the problem.

## **Where is it Used?**

- **Factorial Calculation:** To compute the factorial of a number.
- **Fibonacci Series:** To generate the nth Fibonacci number.
- **Tree Traversal:** In algorithms like Depth-First Search (DFS).
- **Sorting Algorithms:** In algorithms like Quick Sort and Merge Sort.

## **How Does it Work?**

A recursive function typically consists of:
1. **Base Case:** The condition under which the recursion stops.
2. **Recursive Case:** The part of the function where it calls itself with a smaller or simpler input.

## **Python Example**

Here's a simple example of a recursive function to calculate the factorial of a number:

```python
def factorial(n):
    # Base case: factorial of 0 or 1 is 1
    if n == 0 or n == 1:
        return 1
    # Recursive case: n * factorial of (n-1)
    else:
        return n * factorial(n - 1)

# Test the function
print(factorial(5))  # Output: 120
```

- **Explanation:**  
  In this example, the `factorial` function calls itself with `n-1` until it reaches the base case where `n` is 0 or 1. The function then unwinds, multiplying the results together to get the final factorial.
