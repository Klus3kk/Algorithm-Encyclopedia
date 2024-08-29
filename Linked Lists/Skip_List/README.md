# **Skip List**

## **Overview**

A Skip List is a probabilistic data structure that allows for fast search, insertion, and deletion operations by maintaining multiple layers of linked lists.

## **Where is it Used?**

- **Efficient Data Search:** Provides logarithmic time complexity for search operations.
- **Database Indexing:** Used in databases for efficient data access.

## **How Does it Work?**

A Skip List maintains multiple levels of linked lists. Each level is a subset of the level below it, allowing for faster search times by skipping over elements.

## **Python Example**

Implementing a Skip List from scratch can be complex, but here's a simplified concept:

```python
import random

class Node:
    def __init__(self, value, level):
        self.value = value
        self.forward = [None] * (level + 1)

class SkipList:
    def __init__(self, max_level, p):
        self.max_level = max_level
        self.p = p
        self.header = Node(None, max_level)
        self.level = 0

    def random_level(self):
        level = 0
        while random.random() < self.p and level < self.max_level:
            level += 1
        return level

    def insert(self, value):
        update = [None] * (self.max_level + 1)
        current = self.header
        for i in reversed(range(self.level + 1)):
            while current.forward[i] and current.forward[i].value < value:
                current = current.forward[i]
            update[i] = current
        current = current.forward[0]
        if not current or current.value != value:
            new_level = self.random_level()
            if new_level > self.level:
                for i in range(self.level + 1, new_level + 1):
                    update[i] = self.header
                self.level = new_level
            new_node = Node(value, new_level)
            for i in range(new_level + 1):
                new_node.forward[i] = update[i].forward[i]
                update[i].forward[i] = new_node

    def search(self, value):
        current = self.header
        for i in reversed(range(self.level + 1)):
            while current.forward[i] and current.forward[i].value < value:
                current = current.forward[i]
        current = current.forward[0]
        return current and current.value == value

# Example usage
skip_list = SkipList(max_level=3, p=0.5)
skip_list.insert(1)
skip_list.insert(2)
skip_list.insert(3)
print("Search 2:", skip_list.search(2))  # Output: True
print("Search 4:", skip_list.search(4))  # Output: False
```

- **Explanation:**  
  The `SkipList` class implements a basic skip list with insertion and search functionalities, using multiple levels to optimize search time.

