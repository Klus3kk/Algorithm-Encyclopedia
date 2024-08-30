# **Bitwise Operations**

## **Overview:**

Bitwise operations manipulate individual bits of integers. Common operations include AND, OR, XOR, and NOT.

## **Python Example:**

- **Bitwise AND, OR, XOR, and NOT:**

```python
a = 12  # Binary: 1100
b = 7   # Binary: 0111

# Bitwise AND
and_result = a & b  # Binary: 0100, Decimal: 4

# Bitwise OR
or_result = a | b   # Binary: 1111, Decimal: 15

# Bitwise XOR
xor_result = a ^ b  # Binary: 1011, Decimal: 11

# Bitwise NOT
not_result = ~a     # Binary: ...11110011, Decimal: -13 (in 2's complement)

print("AND result:", and_result)
print("OR result:", or_result)
print("XOR result:", xor_result)
print("NOT result:", not_result)
```

## **Explanation:**
- **Bitwise AND (`&`):** Performs a logical AND operation on each bit of two integers.
- **Bitwise OR (`|`):** Performs a logical OR operation on each bit of two integers.
- **Bitwise XOR (`^`):** Performs a logical XOR operation on each bit of two integers.
- **Bitwise NOT (`~`):** Inverts all the bits of the integer.

