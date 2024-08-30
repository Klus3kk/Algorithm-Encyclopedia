# **Permutations and Combinations**

## **Overview:**

Permutations and combinations are methods for counting the number of ways to arrange or select items. Permutations consider the order of arrangement, while combinations do not.

## **Python Example:**

```python
from itertools import permutations, combinations

# Permutations
def generate_permutations(elements):
    return list(permutations(elements))

# Combinations
def generate_combinations(elements, r):
    return list(combinations(elements, r))

# Example usage
elements = [1, 2, 3]
print("Permutations:", generate_permutations(elements))
print("Combinations (r=2):", generate_combinations(elements, 2))
```

## **Explanation:**
- **Permutations:** Generate all possible orders of the elements.
- **Combinations:** Generate all possible selections of elements of a specific length.

