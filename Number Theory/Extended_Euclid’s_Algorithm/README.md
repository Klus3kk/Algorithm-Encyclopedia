# **Extended Euclid’s Algorithm**

## **Overview:**
The Extended Euclid’s Algorithm finds integers \( x \) and \( y \) such that \( ax + by = \text{gcd}(a, b) \). This is useful for solving linear Diophantine equations and in modular arithmetic.

## **Python Code Example:**

```python
def extended_gcd(a, b):
    if a == 0:
        return (b, 0, 1)
    gcd, x1, y1 = extended_gcd(b % a, a)
    x = y1 - (b // a) * x1
    y = x1
    return (gcd, x, y)

# Example usage
a, b = 30, 20
gcd, x, y = extended_gcd(a, b)
print(f"gcd: {gcd}, x: {x}, y: {y}")
```

