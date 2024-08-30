# **Power of Two**

## **Overview:**

Determining if a number is a power of two can be done efficiently using bit manipulation.

## **Python Example:**

- **Check Power of Two:**

```python
def is_power_of_two(n):
    return n > 0 and (n & (n - 1)) == 0

# Example usage
number = 16
print("Is power of two:", is_power_of_two(number))
```

## **Explanation:**
- **Power of Two Check:** A number is a power of two if it has exactly one bit set in its binary representation. `(n & (n - 1))` will be 0 if `n` is a power of two.

