# **Minimum Meeting Rooms**

## **Overview:**

Finding the minimum number of meeting rooms required involves scheduling multiple meetings such that no two meetings overlap in the same room.

## **Python Example:**

- **Find Minimum Meeting Rooms:**

```python
import heapq

def min_meeting_rooms(intervals):
    if not intervals:
        return 0
    
    intervals.sort(key=lambda x: x[0])
    end_times = []
    
    heapq.heappush(end_times, intervals[0][1])
    
    for i in range(1, len(intervals)):
        if intervals[i][0] >= end_times[0]:
            heapq.heappop(end_times)
        heapq.heappush(end_times, intervals[i][1])
    
    return len(end_times)

# Example usage
intervals = [(0, 30), (5, 10), (15, 20)]
print("Minimum number of meeting rooms:", min_meeting_rooms(intervals))
```

## **Explanation:**
- **Minimum Meeting Rooms:** Intervals are sorted by their start times. A min-heap is used to track the end times of ongoing meetings. The size of the heap at any time represents the number of rooms required.

