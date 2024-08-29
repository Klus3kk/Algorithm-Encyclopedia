# **Bloom Filters**

## **Overview**

A Bloom Filter is a probabilistic data structure used to test whether an element is a member of a set. It allows for efficient space utilization but with a possibility of false positives.

## **Where is it Used?**

- **Set Membership Testing:** In applications where space efficiency is crucial and occasional false positives are acceptable.
- **Database Systems:** To quickly check if an element might be present before performing more expensive operations.

## **How Does it Work?**

A Bloom Filter uses multiple hash functions and a bit array. Elements are hashed into several positions in the array, which are set to 1. To check for membership, the same hash functions are used, and if all corresponding bits are set to 1, the element might be present.

## **Python Example**

Here’s a simple implementation of a Bloom Filter:

```python
import hashlib

class BloomFilter:
    def __init__(self, size, hash_count):
        self.size = size
        self.hash_count = hash_count
        self.bit_array = [0] * size

    def _hashes(self, item):
        result = []
        for i in range(self.hash_count):
            hash_value = int(hashlib.md5(item.encode() + str(i).encode()).hexdigest(), 16)
            result.append(hash_value % self.size)
        return result

    def add(self, item):
        for hash_value in self._hashes(item):
            self.bit_array[hash_value] = 1

    def check(self, item):
        return all(self.bit_array[hash_value] for hash_value in self._hashes(item))

# Example usage
bloom_filter = BloomFilter(size=100, hash_count=3)
bloom_filter.add("test")
print("Is 'test' in bloom filter?", bloom_filter.check("test"))  # Output: True
print("Is 'not_test' in bloom filter?", bloom_filter.check("not_test"))  # Output: False or True (due to false positives)
```

- **Explanation:**  
  The `BloomFilter` class uses multiple hash functions to set bits in the bit array. The `check` method verifies if an element might be present, with the possibility of false positives.

