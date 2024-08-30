# **Heap Operations (Insert, Delete, Extract Min/Max)**

## **Overview:**

Key operations on a heap include:
- **Insertion:** Adding an element while maintaining the heap property.
- **Deletion:** Removing an element, typically the root, and restructuring the heap.
- **Extract Min/Max:** Removing and returning the smallest (Min-Heap) or largest (Max-Heap) element.

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

    def delete(self, item):
        # Remove item from heap and maintain heap property
        index = self.heap.index(item)
        self.heap[index] = self.heap[-1]
        self.heap.pop()
        if index < len(self.heap):
            heapq._siftup(self.heap, index)
            heapq._siftdown(self.heap, 0, index)

    def extract_min(self):
        return self.pop()  # Equivalent to pop in Min-Heap

    def extract_max(self):
        return -self.pop()  # Equivalent to pop in Max-Heap

# Example usage
heap = BinaryHeap(min_heap=True)
heap.push(10)
heap.push(5)
heap.push(20)
print("Heap after insertions:", heap.heap)

# Delete and Extract Min
heap.delete(10)
print("Heap after deletion:", heap.heap)
print("Extracted Min:", heap.extract_min())
```

## **Explanation:**
- **Insertion:** Adds elements while preserving heap order.
- **Deletion:** Locates an element, replaces it with the last element, and then restores heap order using `_siftup` and `_siftdown`.
- **Extract Min/Max:** Retrieves and removes the root element.

