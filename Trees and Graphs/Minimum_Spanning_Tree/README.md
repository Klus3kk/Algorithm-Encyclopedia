# **Minimum Spanning Tree**

## **Overview**

A Minimum Spanning Tree (MST) is a subset of the edges of a connected graph that connects all vertices with the minimum total edge weight.

## **Where is it Used?**

- **Network Design:** Designing cost-effective networks.
- **Cluster Analysis:** Grouping nodes into clusters with minimal connection cost.

## **How Does it Work?**

- **Kruskal's Algorithm:** Add edges in increasing order of weight, avoiding cycles.
- **Prim's Algorithm:** Grow the MST one edge at a time from a starting vertex.

## **Python Example**

Here’s an example of Kruskal’s Algorithm for finding an MST:

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

def kruskal(nodes, edges):
    uf = UnionFind(nodes)
    mst = []
    edges = sorted(edges, key=lambda x: x[2])
    for u, v, weight in edges:
        if uf.find(u) != uf.find(v):
            uf.union(u, v)
            mst.append((u, v, weight))
    return mst

# Example usage
nodes = 4
edges = [
    (0, 1, 10),
    (0, 2, 6),
    (0, 3, 5),
    (1, 3, 15),
    (2, 3, 4)
]
print("Minimum Spanning Tree:", kruskal(nodes, edges))
# Output: Minimum Spanning Tree: [(2, 3, 4), (0, 3, 5), (0, 1, 10)]
```

- **Explanation:**  
  Kruskal’s Algorithm uses Union-Find to avoid cycles and find the MST efficiently.

