# **Miller-Rabin Primality Test**

## **Overview:**
The Miller-Rabin primality test is a probabilistic test to determine if a number is prime. It is more efficient than deterministic tests for large numbers.

## **Python Code Example:**

```python
import random

def miller_rabin(n, k=5):
    def is_composite(a, d, n):
        if pow(a, d, n) == 1:
            return False
        while d < n - 1:
            if pow(a, d, n) == n - 1:
                return False
            d *= 2
        return True

    if n in (2, 3):
        return True
    if n < 2 or n % 2 == 0:
        return False
    d = n - 1
    while d % 2 == 0:
        d //= 2
    for _ in range(k):
        a = random.randrange(2, n - 1)
        if is_composite(a, d, n):
            return False
    return True

# Example usage
n = 17
print(miller_rabin(n))
```

