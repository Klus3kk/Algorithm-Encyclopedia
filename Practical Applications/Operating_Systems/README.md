# **Operating Systems**

## **Overview:**

Operating Systems Algorithms handle resource management, scheduling, and process control in computer systems. They ensure efficient execution and management of processes and memory.

## **Subtopics:**

- **Page Replacement Algorithms:** Manage the swapping of pages between RAM and disk.
- **Process Scheduling:** Determines the order in which processes are executed.

## **Python Example (LRU Page Replacement):**

```python
from collections import OrderedDict

class LRUCache:
    def __init__(self, capacity):
        self.cache = OrderedDict()
        self.capacity = capacity
    
    def get(self, key):
        if key not in self.cache:
            return -1
        self.cache.move_to_end(key)
        return self.cache[key]
    
    def put(self, key, value):
        if key in self.cache:
            self.cache.move_to_end(key)
        self.cache[key] = value
        if len(self.cache) > self.capacity:
            self.cache.popitem(last=False)

# Example usage
lru_cache = LRUCache(2)
lru_cache.put(1, 1)
lru_cache.put(2, 2)
print("Get 1:", lru_cache.get(1))
lru_cache.put(3, 3)
print("Get 2:", lru_cache.get(2))
```

## **Explanation:**
- **LRU Cache:** Implements Least Recently Used page replacement by maintaining an ordered dictionary that evicts the least recently used items when the capacity is exceeded.

