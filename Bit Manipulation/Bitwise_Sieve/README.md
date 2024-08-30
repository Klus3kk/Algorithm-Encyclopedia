# **Bitwise Sieve**

## **Overview:**

A bitwise sieve is an efficient way to find all primes up to a given number using bit manipulation to mark non-prime numbers.

## **Python Example:**

- **Sieve of Eratosthenes using Bitwise Operations:**

```python
def bitwise_sieve(limit):
    sieve = bytearray((1,) * (limit + 1))
    sieve[0] = sieve[1] = 0
    p = 2
    while p * p <= limit:
        if sieve[p]:
            for i in range(p * p, limit + 1, p):
                sieve[i] = 0
        p += 1
    return [p for p in range(limit + 1) if sieve[p]]

# Example usage
limit = 30
print("Prime numbers up to", limit, ":", bitwise_sieve(limit))
```

## **Explanation:**
- **Bitwise Sieve:** Uses bitwise operations to efficiently mark non-prime numbers and collect all primes up to a specified limit.

