# **Point in Polygon**

## **Overview:**
Determining whether a point lies inside a polygon is a fundamental geometric problem. Techniques like Ray-casting and the Winding Number method are commonly used for this purpose.

## **Ray-Casting Algorithm:**

The Ray-casting algorithm checks how many times a ray originating from the point crosses the edges of the polygon. An odd number of crossings indicates that the point is inside the polygon.

## **Python Code Example:**

```python
def point_in_polygon(point, polygon):
    def is_point_inside(p, poly):
        x, y = p
        n = len(poly)
        inside = False
        p1x, p1y = poly[0]
        for i in range(n + 1):
            p2x, p2y = poly[i % n]
            if y > min(p1y, p2y):
                if y <= max(p1y

, p2y):
                    if x <= max(p1x, p2x):
                        if p1y != p2y:
                            xints = (y - p1y) * (p2x - p1x) / (p2y - p1y) + p1x
                        if p1x == p2x or x <= xints:
                            inside = not inside
            p1x, p1y = p2x, p2y
        return inside

    return is_point_inside(point, polygon)

# Example usage
polygon = [(0, 0), (1, 0), (1, 1), (0, 1)]
point = (0.5, 0.5)
print(point_in_polygon(point, polygon))
```

