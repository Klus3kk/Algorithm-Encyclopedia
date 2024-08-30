# **Sweep Line Algorithm**

## **Overview:**

The sweep line algorithm is used to solve various problems involving intervals, such as finding overlapping intervals, by "sweeping" a line across the data.

## **Python Example:**

- **Finding All Overlapping Intervals Using Sweep Line:**

```python
def find_overlaps(intervals):
    events = []
    for start, end in intervals:
        events.append((start, 'start'))
        events.append((end, 'end'))
    
    events.sort(key=lambda x: (x[0], x[1] == 'start'))
    active_intervals = 0
    overlaps = []
    
    for time, event_type in events:
        if event_type == 'start':
            active_intervals += 1
        else:
            active_intervals -= 1
        if active_intervals > 1:
            overlaps.append(time)
    
    return overlaps

# Example usage
intervals = [(1, 5), (2, 6), (8, 10), (7, 9)]
print("Times with overlaps:", find_overlaps(intervals))
```

## **Explanation:**
- **Sweep Line:** The algorithm processes start and end events of intervals, sorting them to determine the times when overlaps occur. It tracks the number of active intervals to detect overlaps.

