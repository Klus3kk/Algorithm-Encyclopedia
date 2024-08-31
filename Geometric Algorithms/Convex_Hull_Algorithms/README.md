# **Convex Hull Algorithms**

## **Overview:**
Convex hull algorithms compute the convex hull of a set of points. The convex hull is the smallest convex boundary that encloses all the points in a given set. This concept is fundamental in computational geometry, computer graphics, and geographic information systems.

## **Quickhull Algorithm:**

Quickhull is a popular algorithm for computing the convex hull. It uses a divide-and-conquer approach to recursively find the extreme points of the set.

## **Python Code Example:**

```python
def quickhull(points):
    def hull(points, p1, p2):
        if not points:
            return []
        left = [p for p in points if cross_product(p1, p2, p) > 0]
        right = [p for p in points if cross_product(p1, p2, p) < 0]
        if not left:
            return [p1, p2]
        farthest = max(left, key=lambda p: distance(p1, p2, p))
        left.remove(farthest)
        return hull(left, p1, farthest) + hull(right, farthest, p2)

    def cross_product(p1, p2, p3):
        return (p2[0] - p1[0]) * (p3[1] - p1[1]) - (p2[1] - p1[1]) * (p3[0] - p1[0])

    def distance(p1, p2, p):
        return abs(cross_product(p1, p2, p)) / ((p2[0] - p1[0]) ** 2 + (p2[1] - p1[1]) ** 2) ** 0.5

    points = sorted(set(points))
    if len(points) <= 1:
        return points
    leftmost = min(points)
    rightmost = max(points)
    return hull(points, leftmost, rightmost) + hull(points, rightmost, leftmost)

# Example usage
points = [(0, 0), (1, 1), (2, 2), (2, 0), (0, 2)]
print(quickhull(points))
```

