# **Union-Find**

## **Overview:**

Union-Find (Disjoint Set Union) is a data structure for managing a partition of a set into disjoint subsets, supporting union and find operations efficiently.

## **Python Example:**

- **Union-Find Data Structure:**

```python
class UnionFind:
    def __init__(self, size):
        self.parent = list(range(size))
        self.rank = [0] * size
    
    def find(self, x):
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]
    
    def union(self, x, y):
        rootX = self.find(x)
        rootY = self.find(y)
        
        if rootX != rootY:
            if self.rank[rootX] > self.rank[rootY]:
                self.parent[rootY] = rootX
            elif self.rank[rootX] < self.rank[rootY]:
                self.parent[rootX] = rootY
            else:
                self.parent[rootY] = rootX
                self.rank[rootX] += 1

# Example usage
uf = UnionFind(5)
uf.union(0, 1)
uf.union(1, 2)
print("Find(2):", uf.find(2))
print("Find(0):", uf.find(0))
```

## **Explanation:**
- **Union-Find:** Manages disjoint sets with path compression in `find` and union by rank to keep the tree flat, improving efficiency.

