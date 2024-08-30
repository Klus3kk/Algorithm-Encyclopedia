# **Counting Set Bits**

## **Overview:**

Counting set bits involves determining the number of bits set to 1 in the binary representation of an integer.

## **Python Example:**

- **Count Set Bits:**

```python
def count_set_bits(n):
    count = 0
    while n:
        count += n & 1
        n >>= 1
    return count

# Example usage
number = 29  # Binary: 11101
print("Number of set bits:", count_set_bits(number))
```

## **Explanation:**
- **Count Set Bits:** Iteratively counts the number of 1s in the binary representation of a number.

