# **Kruskal’s MST Algorithm**

## **Overview:**

Kruskal’s Algorithm finds the Minimum Spanning Tree (MST) of a graph by sorting edges in increasing order of weight and adding edges to the MST if they don’t form a cycle.

## **Python Example:**

```python
class UnionFind:
    def __init__(self, size):
        self.parent = list(range(size))
        self.rank = [0] * size

    def find(self, u):
        if self.parent[u] != u:
            self.parent[u] = self.find(self.parent[u])
        return self.parent[u]

    def union(self, u, v):
        root_u = self.find(u)
        root_v = self.find(v)
        if root_u != root_v:
            if self.rank[root_u] > self.rank[root_v]:
                self.parent[root_v] = root_u
            elif self.rank[root_u] < self.rank[root_v]:
                self.parent[root_u] = root_v
            else:
                self.parent[root_v] = root_u
                self.rank[root_u] += 1

def kruskal(vertices, edges):
    edges.sort(key=lambda x: x[2])  # Sort edges by weight
    uf = UnionFind(vertices)
    mst = []
    for u, v, weight in edges:
        if uf.find(u) != uf.find(v):
            uf.union(u, v)
            mst.append((u, v, weight))
    return mst

# Example usage
vertices = 4
edges = [
    (0, 1, 10),
    (0, 2, 6),
    (0, 3, 5),
    (1, 3, 15),
    (2, 3, 4)
]
print("Kruskal's MST:", kruskal(vertices, edges))
```

## **Explanation:**
- **Sort Edges:** Edges are sorted by weight.
- **Union-Find:** Used to detect and avoid cycles.
- **MST Construction:** Edges are added to the MST if they don’t form a cycle.
