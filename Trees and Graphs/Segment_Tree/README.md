# **Segment Tree**

## **Overview**

A Segment Tree is a data structure that allows for efficient range queries and updates on an array.

## **Where is it Used?**

- **Range Queries:** Queries like sum, minimum, and maximum over a range.
- **Dynamic Updates:** Handling updates to the array while maintaining query efficiency.

## **How Does it Work?**

- **Build:** Construct the segment tree from the array.
- **Query:** Retrieve information for a specified range.
- **Update:** Modify the value at a specific index.

## **Python Example**

Here’s an example of a Segment Tree for range sum queries:

```python
class SegmentTree:
    def __init__(self, data):
        self.n = len(data)
        self.tree = [0] * (2 * self.n)
        self.build(data)

    def build(self, data):
        for i in range(self.n):
            self.tree[self.n + i] = data[i]
        for i in range(self.n - 1, 0, -1):
            self.tree[i] = self.tree[2 * i] + self.tree[2 * i + 1]

    def query(self, left, right):
        result = 0
        left +=

 self.n
        right += self.n + 1
        while left < right:
            if left % 2:
                result += self.tree[left]
                left += 1
            if right % 2:
                right -= 1
                result += self.tree[right]
            left //= 2
            right //= 2
        return result

# Example usage
data = [1, 3, 5, 7, 9, 11]
seg_tree = SegmentTree(data)
print("Sum of range [1, 4]:", seg_tree.query(1, 4))  # Output: 24
```

- **Explanation:**  
  The `SegmentTree` class handles range sum queries efficiently with a segment tree structure.

