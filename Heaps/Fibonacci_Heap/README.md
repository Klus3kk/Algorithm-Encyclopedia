# **Fibonacci Heap**

## **Overview:**

A Fibonacci Heap is a more advanced heap structure that allows for more efficient amortized time complexities for various operations. It consists of a collection of heap-ordered trees. Key operations include insertion, deletion, and decrease-key.

## **Python Example:**

```python
class FibonacciHeapNode:
    def __init__(self, key):
        self.key = key
        self.degree = 0
        self.child = None
        self.next = self
        self.prev = self
        self.mark = False

class FibonacciHeap:
    def __init__(self):
        self.min_node = None
        self.total_nodes = 0

    def insert(self, key):
        new_node = FibonacciHeapNode(key)
        if self.min_node is None:
            self.min_node = new_node
        else:
            self._link(self.min_node, new_node)
            if key < self.min_node.key:
                self.min_node = new_node
        self.total_nodes += 1

    def _link(self, min_node, new_node):
        # Link new_node into the circular doubly linked list of min_node
        new_node.next = min_node.next
        min_node.next.prev = new_node
        min_node.next = new_node
        new_node.prev = min_node

# Example usage
f_heap = FibonacciHeap()
f_heap.insert(10)
f_heap.insert(20)
f_heap.insert(5)
print("Fibonacci Heap min key:", f_heap.min_node.key)
```

## **Explanation:**
- **Insert:** Adds a new node and updates the minimum node if necessary.
- **Link:** Inserts a node into the circular doubly linked list.

