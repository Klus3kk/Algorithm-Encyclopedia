# **Modular Inverse**

## **Overview:**
The modular inverse of a number \( a \) modulo \( m \) is a number \( x \) such that \( ax \equiv 1 \ (\text{mod} \ m) \). It is crucial in solving modular equations and is widely used in cryptography.

## **Python Code Example:**

```python
def modular_inverse(a, m):
    def extended_gcd(a, b):
        if a == 0:
            return (b, 0, 1)
        gcd, x1, y1 = extended_gcd(b % a, a)
        x = y1 - (b // a) * x1
        y = x1
        return (gcd, x, y)

    gcd, x, y = extended_gcd(a, m)
    if gcd != 1:
        raise ValueError("No modular inverse exists")
    return x % m

# Example usage
a, m = 3, 11
print(modular_inverse(a, m))
```

