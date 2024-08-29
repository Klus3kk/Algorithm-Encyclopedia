# **Collision Handling (Chaining, Open Addressing)**

## **Overview**

Collision handling techniques are used to manage situations where multiple keys hash to the same index in a hashmap or set. Two common techniques are chaining and open addressing.

## **Where is it Used?**

- **Hashmaps and Sets:** To handle key collisions efficiently.
- **Database Indexing:** Managing collisions in hash-based indexes.

## **How Does it Work?**

- **Chaining:** Uses a list (or another data structure) to store all elements that hash to the same index. Each index in the hash table points to a list of entries.
  
- **Open Addressing:** Probes for alternative positions in the hash table when a collision occurs. Common methods include linear probing, quadratic probing, and double hashing.

## **Python Example**

Here’s a simple implementation of collision handling using chaining in a hashmap:

```python
class HashMap:
    def __init__(self, size):
        self.size = size
        self.table = [[] for _ in range(size)]

    def hash_function(self, key):
        return hash(key) % self.size

    def insert(self, key, value):
        index = self.hash_function(key)
        for pair in self.table[index]:
            if pair[0] == key:
                pair[1] = value
                return
        self.table[index].append([key, value])

    def get(self, key):
        index = self.hash_function(key)
        for pair in self.table[index]:
            if pair[0] == key:
                return pair[1]
        return None

# Example usage
hash_map = HashMap(10)
hash_map.insert("key1", "value1")
print("Value for 'key1':", hash_map.get("key1"))  # Output: value1
```

- **Explanation:**  
  The `HashMap` class uses chaining to handle collisions. The `insert` method adds key-value pairs, and the `get` method retrieves values based on the key.

