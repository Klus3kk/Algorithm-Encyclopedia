# **Queue Operations**

## **Overview**

A queue is a linear data structure that follows the First-In-First-Out (FIFO) principle. Elements are added at the rear and removed from the front.

## **Where is it Used?**

- **Scheduling:** Managing tasks or processes in order.
- **Breadth-First Search:** Implementing BFS in graphs.
- **Buffer Management:** Managing data streams and IO operations.

## **How Does it Work?**

- **Enqueue:** Add an element to the rear of the queue.
- **Dequeue:** Remove the element from the front of the queue.
- **Peek:** Retrieve the front element without removing it.

## **Python Example**

Here’s an implementation of a queue using `collections.deque`:

```python
from collections import deque

class Queue:
    def __init__(self):
        self.queue = deque()

    def enqueue(self, item):
        self.queue.append(item)

    def dequeue(self):
        if not self.is_empty():
            return self.queue.popleft()
        raise IndexError("Dequeue from empty queue")

    def peek(self):
        if not self.is_empty():
            return self.queue[0]
        raise IndexError("Peek from empty queue")

    def is_empty(self):
        return len(self.queue) == 0

# Example usage
queue = Queue()
queue.enqueue(1)
queue.enqueue(2)
print("Peek:", queue.peek())  # Output: 1
print("Dequeue:", queue.dequeue())  # Output: 1
print("Dequeue:", queue.dequeue())  # Output: 2
```

- **Explanation:**  
  The `Queue` class uses `deque` for efficient enqueue and dequeue operations and provides methods to peek and check if the queue is empty.

