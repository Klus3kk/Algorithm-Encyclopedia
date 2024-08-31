# **Database Algorithms**

## **Overview:**

Database Algorithms are used to efficiently manage and query data in databases. They include indexing methods and query optimization techniques that enhance performance.

## **Subtopics:**

- **B-Trees:** Self-balancing tree data structures used for indexing.
- **Hash Indexing:** Uses hash functions to index data for quick retrieval.

## **Python Example (B-Tree):**

```python
class BTreeNode:
    def __init__(self, t, leaf=False):
        self.t = t
        self.leaf = leaf
        self.keys = []
        self.children = []

class BTree:
    def __init__(self, t):
        self.root = BTreeNode(t, leaf=True)
        self.t = t
    
    def traverse(self, node=None):
        if node is None:
            node = self.root
        for key in node.keys:
            print(key, end=' ')
        if not node.leaf:
            for child in node.children:
                self.traverse(child)
                
# Example usage
b_tree = BTree(t=3)
b_tree.root.keys = [10, 20]
b_tree.root.children = [BTreeNode(3, leaf=True), BTreeNode(3, leaf=True)]
b_tree.traverse()
```

## **Explanation:**
- **B-Trees:** Maintain sorted data and allow searches, insertions, and deletions in logarithmic time, making them suitable for database indexing.

## **Python Example (Hash Indexing):**

```python
class HashIndex:
    def __init__(self):
        self.table = {}

    def insert(self, key, value):
        self.table[key] = value

    def search(self, key):
        return self.table.get(key, None)

# Example usage
hash_index = HashIndex()
hash_index.insert('name', 'Alice')
print("Search for 'name':", hash_index.search('name'))
```

## **Explanation:**
- **Hash Indexing:** Uses a hash function to index data, allowing for average-case constant time complexity for insertions and lookups.

