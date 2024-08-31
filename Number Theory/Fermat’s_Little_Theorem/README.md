# **Fermat’s Little Theorem**

## **Overview:**
Fermat’s Little Theorem states that if \( p \) is a prime and \( a \) is not divisible by \( p \), then \( a^{p-1} \equiv 1 \ (\text{mod} \ p) \). It is used in primality testing and cryptographic algorithms.

## **Python Code Example:**

```python
def fermat_little_theorem(a, p):
    if pow(a, p - 1, p) == 1:
        return True
    return False

# Example usage
a, p = 2, 7
print(fermat_little_theorem(a, p))
```

