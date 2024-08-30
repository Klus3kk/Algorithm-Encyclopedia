# **Bit Masks**

## **Overview:**

Bit masks are used to manipulate specific bits within an integer. They allow checking, setting, or clearing individual bits.

## **Python Example:**

- **Checking, Setting, and Clearing Bits:**

```python
def check_bit(number, position):
    return (number & (1 << position)) != 0

def set_bit(number, position):
    return number | (1 << position)

def clear_bit(number, position):
    return number & ~(1 << position)

number = 0b1010  # Binary: 1010

# Check if the 1st bit is set
print("Bit at position 1 is set:", check_bit(number, 1))

# Set the 2nd bit
number = set_bit(number, 2)
print("After setting the 2nd bit:", bin(number))

# Clear the 3rd bit
number = clear_bit(number, 3)
print("After clearing the 3rd bit:", bin(number))
```

## **Explanation:**
- **Checking Bit:** Uses a mask to check if a specific bit is set.
- **Setting Bit:** Uses a mask to set a specific bit to 1.
- **Clearing Bit:** Uses a mask to clear a specific bit to 0.

