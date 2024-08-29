# **Circular Queue**

## **Overview**

A circular queue is a type of queue where the end of the queue connects back to the beginning, forming a circle. This helps in efficient utilization of space.

## **Where is it Used?**

- **Buffer Management:** Implementing circular buffers in systems.
- **Round-Robin Scheduling:** Managing tasks in a cyclic order.

## **How Does it Work?**

- **Enqueue and Dequeue:** Use modulo arithmetic to wrap around the end of the queue.

## **Python Example**

Here’s a simple implementation of a circular queue:

```python
class CircularQueue:
    def __init__(self, size):
        self.size = size
        self.queue = [None] * size
        self.front = 0
        self.rear = 0

    def enqueue(self, item):
        if (self.rear + 1) % self.size == self.front:
            raise OverflowError("Queue is full")
        self.queue[self.rear] = item
        self.rear = (self.rear + 1) % self.size

    def dequeue(self):
        if self.front == self.rear:
            raise IndexError("Queue is empty")
        item = self.queue[self.front]
        self.front = (self.front + 1) % self.size
        return item

    def peek(self):
        if self.front == self.rear:
            raise IndexError("Queue is empty")
        return self.queue[self.front]

    def is_empty(self):
        return self.front == self.rear

# Example usage
cq = CircularQueue(3)
cq.enqueue(1)
cq.enqueue(2)
print("Dequeue:", cq.dequeue())  # Output: 1
cq.enqueue(3)
cq.enqueue(4)
print("Peek:", cq.peek())  # Output: 2
```

- **Explanation:**  
  The `CircularQueue` class maintains the queue in a circular fashion, ensuring efficient space utilization.
