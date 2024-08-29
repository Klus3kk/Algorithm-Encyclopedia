# **Recurrence Relations**

## **Overview**

Recurrence relations are equations that define sequences recursively. They are used to express the time complexity of recursive algorithms and to solve problems in dynamic programming.

## **Where is it Used?**

- **Algorithm Analysis:** To determine the time complexity of recursive algorithms.
- **Dynamic Programming:** To solve optimization problems.
- **Combinatorics:** In counting problems where each solution builds on smaller solutions.

## **How Does it Work?**

A recurrence relation expresses the value of a function or sequence at any point `n` in terms of its values at earlier points. For example:
- **T(n) = 2T(n/2) + n:** This is a common recurrence relation that arises in the analysis of the Merge Sort algorithm.

## **Python Example**

Let's consider the recurrence relation for calculating Fibonacci numbers:

```python
def fibonacci(n):
    # Base cases
    if n == 0:
        return 0
    elif n == 1:
        return 1
    # Recursive case
    else:
        return fibonacci(n - 1) + fibonacci(n - 2)

# Test the function
print(fibonacci(6))  # Output: 8
```

- **Explanation:**  
  The Fibonacci sequence is defined by the recurrence relation `F(n) = F(n-1) + F(n-2)`, with base cases `F(0) = 0` and `F(1) = 1`. The function computes the value of `F(n)` by recursively summing the two preceding values.
