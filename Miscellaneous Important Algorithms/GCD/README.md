# **Euclid’s Algorithm (GCD)**

## **Overview:**

Euclid’s Algorithm computes the Greatest Common Divisor (GCD) of two integers. The GCD is the largest number that divides both integers without leaving a remainder.

## **Python Example:**

```python
def euclidean_gcd(a, b):
    while b:
        a, b = b, a % b
    return a

# Example usage
print("GCD of 48 and 18:", euclidean_gcd(48, 18))
```

## **Explanation:**
- **Euclid’s Algorithm:** Iteratively replaces the larger number with its remainder when divided by the smaller number until the remainder is zero. The non-zero divisor at this point is the GCD.

