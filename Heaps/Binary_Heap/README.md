# **Binary Heap (Min-Heap, Max-Heap)**

## **Overview:**

A Binary Heap is a complete binary tree that satisfies the heap property:
- **Min-Heap:** In a Min-Heap, the value at the root is the smallest, and each parent node is less than or equal to its children.
- **Max-Heap:** In a Max-Heap, the value at the root is the largest, and each parent node is greater than or equal to its children.

## **Python Example:**

```python
import heapq

class BinaryHeap:
    def __init__(self, min_heap=True):
        self.heap = []
        self.min_heap = min_heap

    def push(self, item):
        if not self.min_heap:
            item = -item
        heapq.heappush(self.heap, item)

    def pop(self):
        item = heapq.heappop(self.heap)
        if not self.min_heap:
            item = -item
        return item

    def peek(self):
        item = self.heap[0]
        if not self.min_heap:
            item = -item
        return item

    def __len__(self):
        return len(self.heap)

# Example usage
min_heap = BinaryHeap(min_heap=True)
max_heap = BinaryHeap(min_heap=False)

# Min-Heap Operations
min_heap.push(5)
min_heap.push(3)
min_heap.push(8)
print("Min-Heap:", [min_heap.pop() for _ in range(len(min_heap))])

# Max-Heap Operations
max_heap.push(5)
max_heap.push(3)
max_heap.push(8)
print("Max-Heap:", [max_heap.pop() for _ in range(len(max_heap))])
```

## **Explanation:**
- **Min-Heap:** The `min_heap` parameter controls whether the heap is a Min-Heap or Max-Heap. 
- **Insertion:** Uses `heapq.heappush` to maintain heap order.
- **Deletion:** Uses `heapq.heappop` to remove and return the smallest (or largest) element while preserving the heap property.
