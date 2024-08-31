# **Exponentiation by Squaring**

## **Overview:**
Exponentiation by squaring is an efficient algorithm for computing large powers of numbers modulo some value. It reduces the number of multiplicative operations needed.

## **Python Code Example:**

```python
def exponentiation_by_squaring(base, exponent, mod):
    result = 1
    base = base % mod
    while exponent > 0:
        if (exponent % 2) == 1:
            result = (result * base) % mod
        exponent = exponent >> 1
        base = (base * base) % mod
    return result

# Example usage
base, exponent, mod = 2, 10, 1000
print(exponentiation_by_squaring(base, exponent, mod))
```

