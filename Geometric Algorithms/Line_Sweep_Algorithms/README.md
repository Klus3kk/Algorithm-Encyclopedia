# **Line Sweep Algorithms**

## **Overview:**
Line sweep algorithms are used to detect intersections between line segments by sweeping a vertical line across the plane. This technique is efficient for problems involving multiple line segments and their intersections.

## **Bentley-Ottmann Algorithm:**

The Bentley-Ottmann algorithm is a line sweep algorithm that detects all intersections among a set of line segments. It maintains an event queue and a status structure to handle intersections.

## **Python Code Example:**

```python
from sortedcontainers import SortedList

def bentley_ottmann(segments):
    events = SortedList()
    for seg in segments:
        events.add((min(seg[0][0], seg[1][0]), seg))
        events.add((max(seg[0][0], seg[1][0]), seg))

    def intersects(seg1, seg2):
        def ccw(p1, p2, p3):
            return (p3[1] - p1[1]) * (p2[0] - p1[0]) > (p2[1] - p1[1]) * (p3[0] - p1[0])
        return ccw(seg1[0], seg2[0], seg2[1]) != ccw(seg1[1], seg2[0], seg2[1]) and \
               ccw(seg1[0], seg1[1], seg2[0]) != ccw(seg1[0], seg1[1], seg2[1])

    active_segments = SortedList()
    intersections = []
    while events:
        event = events.pop(0)
        x, seg = event
        if seg in active_segments:
            active_segments.remove(seg)
            for other in active_segments:
                if intersects(seg, other):
                    intersections.append((seg, other))
        else:
            active_segments.add(seg)
            for other in active_segments:
                if intersects(seg, other):
                    intersections.append((seg, other))
    return intersections

# Example usage
segments = [((0, 0), (2, 2)), ((0, 2), (2, 0))]
print(bentley_ottmann(segments))
```

