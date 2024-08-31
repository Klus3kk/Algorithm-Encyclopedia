# **Convex Hull Algorithms**

## **Overview:**

Convex Hull algorithms find the smallest convex polygon that can enclose a set of points in a plane. Important algorithms include Graham Scan and Jarvis March.

## **Python Example (Graham Scan):**

```python
from functools import cmp_to_key

def graham_scan(points):
    def orientation(p, q, r):
        return (q[1] - p[1]) * (r[0] - q[0]) - (q[0] - p[0]) * (r[1] - q[1])
    
    points = sorted(points)
    lower = []
    for p in points:
        while len(lower) >= 2 and orientation(lower[-2], lower[-1], p) <= 0:
            lower.pop()
        lower.append(p)
    
    upper = []
    for p in reversed(points):
        while len(upper) >= 2 and orientation(upper[-2], upper[-1], p) <= 0:
            upper.pop()
        upper.append(p)
    
    return lower[:-1] + upper[:-1]

# Example usage
points = [(0, 0), (2, 2), (1, 1), (2, 1), (3, 0), (3, 3)]
print("Convex Hull:", graham_scan(points))
```

## **Explanation:**
- **Graham Scan:** Sorts the points and constructs the convex hull by processing points in increasing order and ensuring the hull is convex using the orientation test.

