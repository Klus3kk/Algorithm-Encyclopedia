# **Number Theory**

## **Overview:**

Number Theory deals with properties and relationships of integers. Key algorithms include modular arithmetic, prime testing, and factoring techniques.

## **Subtopics:**

- **Modular Arithmetic:** Operations involving remainders.
- **Chinese Remainder Theorem:** Solves systems of congruences.
- **Sieve of Eratosthenes:** Finds all primes up to a given limit.

## **Python Example (Sieve of Eratosthenes):**

```python
def sieve_of_eratosthenes(n):
    is_prime = [True] * (n + 1)
    p = 2
    while (p * p <= n):
        if (is_prime[p] == True):
            for i in range(p * p, n + 1, p):
                is_prime[i] = False
        p += 1
    return [p for p in range(2, n) if is_prime[p]]

# Example usage
print("Primes up to 30:", sieve_of_eratosthenes(30))
```

## **Explanation:**
- **Sieve of Eratosthenes:** Efficiently finds all prime numbers up to a given number by iteratively marking multiples of each prime starting from 2.

