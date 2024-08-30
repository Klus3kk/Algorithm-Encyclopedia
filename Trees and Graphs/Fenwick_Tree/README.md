# **Fenwick Tree (Binary Indexed Tree)**

## **Overview**

A Fenwick Tree (or Binary Indexed Tree) is a data structure that provides efficient methods for cumulative frequency tables.

## **Where is it Used?**

- **Prefix Sum Queries:** Calculating prefix sums efficiently.
- **Dynamic Updates:** Updating values while maintaining prefix sum accuracy.

## **How Does it Work?**

- **Update:** Add a value to an element and propagate the change.
- **Query:** Compute the sum of elements from the start up to a specified index.

## **Python Example**

Here’s an example of a Fenwick Tree:

```python
class FenwickTree:
    def __init__(self, size):
        self.size = size
        self.tree = [0] * (size + 1)

    def update(self, index, delta):
        while index <= self.size:
            self.tree[index] += delta
            index += index & -index

    def query(self, index):
        sum_ = 0
        while index > 0:
            sum_ += self.tree[index]
            index -= index & -index
        return sum_

# Example usage
fenwick_tree = FenwickTree(6)
fenwick_tree.update(1, 1)
fenwick_tree.update(2, 3)
fenwick_tree.update(3, 5)
print("Prefix sum up to index 3:", fenwick_tree.query(3))  # Output: 9
```

- **Explanation:**  
  The `FenwickTree` class supports efficient updates and prefix sum queries.

