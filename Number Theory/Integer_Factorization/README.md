# **Integer Factorization**

## **Overview:**
Integer factorization involves breaking down a number into its prime factors. It is a fundamental problem in number theory with applications in cryptography.

## **Pollard's Rho Algorithm:**

Pollard's Rho Algorithm is a method for integer factorization based on random number generation and the use of the greatest common divisor.

## **Python Code Example:**

```python
import math
import random

def pollards_rho(n):
    def gcd(a, b):
        while b:
            a, b = b, a % b
        return a

    def f(x, n):
        return (x * x + 1) % n

    if n % 2 == 0:
        return 2
    x = random.randrange(2, n)
    y = x
    d = 1
    while d == 1:
        x = f(x, n)
        y = f(f(y, n), n)
        d = gcd(abs(x - y), n)
    return d

# Example usage
n = 91
print(pollards_rho(n))
```

