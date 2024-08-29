# **Tail Recursion**

## **Overview**

Tail recursion is a special case of recursion where the recursive call is the last operation in the function. This allows the compiler or interpreter to optimize the recursion, reusing the current function's stack frame for the next recursive call, thus preventing stack overflow in cases of deep recursion.

## **Where is it Used?**

- **Optimized Recursive Functions:** To prevent stack overflow in languages or environments that support tail call optimization.
- **Functional Programming:** Often used in languages like Scheme, Haskell, and Lisp.

## **How Does it Work?**

In tail recursion, the result of the recursive call is immediately returned, with no additional computation after the return. This allows the optimization where the current function's frame is replaced with the next one.

## **Python Example**

Here's an example of a tail-recursive function to calculate the factorial of a number:

```python
def tail_recursive_factorial(n, accumulator=1):
    # Base case
    if n == 0:
        return accumulator
    # Tail recursion: call the function with updated accumulator
    else:
        return tail_recursive_factorial(n - 1, n * accumulator)

# Test the function
print(tail_recursive_factorial(5))  # Output: 120
```

- **Explanation:**  
  The `tail_recursive_factorial` function uses an accumulator to hold the result of the computation. Each recursive call updates the accumulator with the current value of `n * accumulator`, and when `n` reaches 0, the accumulator holds the final result.
