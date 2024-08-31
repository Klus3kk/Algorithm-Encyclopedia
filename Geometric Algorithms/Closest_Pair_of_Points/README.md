# **Closest Pair of Points**

## **Overview:**
The closest pair of points problem involves finding the pair of points in a set with the smallest distance between them. It is a classic problem in computational geometry with applications in clustering and nearest-neighbor searches.

## **Divide and Conquer Approach:**

This algorithm divides the points into two halves, recursively finds the closest pair in each half, and then finds the closest pair that spans both halves.

## **Python Code Example:**

```python
def closest_pair(points):
    def closest_pair_rec(points):
        if len(points) <= 3:
            return min((distance(p1, p2) for p1 in points for p2 in points if p1 != p2), default=float('inf'))
        mid = len(points) // 2
        left = points[:mid]
        right = points[mid:]
        d = min(closest_pair_rec(left), closest_pair_rec(right))
        strip = [p for p in points if abs(p[0] - points[mid][0]) < d]
        strip.sort(key=lambda p: p[1])
        for i in range(len(strip)):
            for j in range(i+1, len(strip)):
                if strip[j][1] - strip[i][1] >= d:
                    break
                d = min(d, distance(strip[i], strip[j]))
        return d

    def distance(p1, p2):
        return ((p1[0] - p2[0]) ** 2 + (p1[1] - p2[1]) ** 2) ** 0.5

    points = sorted(points)
    return closest_pair_rec(points)

# Example usage
points = [(0, 0), (1, 1), (2, 2), (3, 3)]
print(closest_pair(points))
```

