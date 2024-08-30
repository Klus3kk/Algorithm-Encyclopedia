# **Merging Intervals**

## **Overview:**

Merging intervals involves combining overlapping intervals into a single interval.

## **Python Example:**

- **Merge Overlapping Intervals:**

```python
def merge_intervals(intervals):
    if not intervals:
        return []
    
    intervals.sort(key=lambda x: x[0])
    merged = [intervals[0]]
    
    for current in intervals[1:]:
        last = merged[-1]
        if current[0] <= last[1]:  # Overlapping intervals
            last[1] = max(last[1], current[1])
        else:
            merged.append(current)
    
    return merged

# Example usage
intervals = [(1, 3), (2, 6), (8, 10), (15, 18)]
print("Merged intervals:", merge_intervals(intervals))
```

## **Explanation:**
- **Merge Intervals:** Intervals are first sorted by their start times. Overlapping intervals are then merged into a single interval by updating the end time of the last merged interval.

