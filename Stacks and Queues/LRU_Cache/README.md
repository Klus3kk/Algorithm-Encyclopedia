# **LRU Cache**

## **Overview**

The Least Recently Used (LRU) cache is a data structure that maintains a limited number of items, evicting the least recently used item when the cache is full.

## **Where is it Used?**

- **Caching Mechanisms:** Optimizing memory and disk usage by caching frequently accessed data.
- **Database Query Caching:** Improving query performance by caching results.

## **How Does it Work?**

- **Eviction Policy:** Use a combination of a doubly linked list and a hash map to efficiently manage access and eviction.

## **Python Example**

Here’s an implementation of an LRU cache:

```python
from collections import OrderedDict

class LRUCache:
    def __init__(self, capacity):
        self.cache

 = OrderedDict()
        self.capacity = capacity

    def get(self, key):
        if key not in self.cache:
            return -1
        else:
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
print("Get 1:", lru_cache.get(1))  # Output: 1
lru_cache.put(3, 3)  # Evicts key 2
print("Get 2:", lru_cache.get(2))  # Output: -1
lru_cache.put(4, 4)  # Evicts key 1
print("Get 1:", lru_cache.get(1))  # Output: -1
print("Get 3:", lru_cache.get(3))  # Output: 3
print("Get 4:", lru_cache.get(4))  # Output: 4
```

- **Explanation:**  
  The `LRUCache` class uses an `OrderedDict` to manage cache entries and their access order.

