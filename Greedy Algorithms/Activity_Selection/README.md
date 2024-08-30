# **Activity Selection**

## **Overview:**

The Activity Selection Problem involves selecting the maximum number of activities that don't overlap, given their start and end times. It's a classic problem where a greedy approach is optimal.

## **Python Example:**

```python
def activity_selection(start_times, end_times):
    activities = list(zip(start_times, end_times))
    activities.sort(key=lambda x: x[1])  # Sort activities by end time

    selected_activities = []
    last_end_time = -1

    for start, end in activities:
        if start > last_end_time:
            selected_activities.append((start, end))
            last_end_time = end

    return selected_activities

# Example usage
start_times = [1, 3, 0, 5, 8, 5]
end_times = [2, 4, 6, 7, 9, 9]
print("Selected Activities:", activity_selection(start_times, end_times))
```

## **Explanation:**
- **Sort Activities:** Activities are sorted by their end times to maximize the number of non-overlapping activities.
- **Selection:** Activities are chosen if their start time is greater than the end time of the last selected activity.

