# **Interval Overlap**

## **Overview:**

Interval overlap problems involve determining if two or more intervals intersect or overlap.

## **Python Example:**

- **Check Overlap Between Two Intervals:**

```python
def is_overlap(interval1, interval2):
    return not (interval1[1] < interval2[0] or interval2[1] < interval1[0])

# Example usage
interval1 = (1, 5)
interval2 = (4, 8)
print("Intervals overlap:", is_overlap(interval1, interval2))
```

## **Explanation:**
- **Interval Overlap Check:** Two intervals overlap if one interval starts before the other one ends and vice versa.

