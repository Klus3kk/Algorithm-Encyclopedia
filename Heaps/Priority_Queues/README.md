# **Priority Queues**

## **Overview:**

A Priority Queue is an abstract data type that supports operations to insert elements and retrieve the element with the highest (or lowest) priority efficiently. Heaps are often used to implement priority queues.

## **Python Example:**

```python
import heapq

class PriorityQueue:
    def __init__(self):
        self.heap = []

    def push(self, item, priority):
        heapq.heappush(self.heap, (priority, item))

    def pop(self):
        return heapq.heappop(self.heap)[1]

    def peek(self):
        return self.heap[0][1]

# Example usage
pq = PriorityQueue()
pq.push("task1", 2)
pq.push("task2", 1)
pq.push("task3", 3)
print("Peek:", pq.peek())
print("Pop:", pq.pop())
print("Pop:", pq.pop())
```

## **Explanation:**
- **Insertion:** Uses a tuple `(priority, item)` to maintain the priority order.
- **Retrieve Highest Priority:** Retrieves and removes the element with the highest priority.

