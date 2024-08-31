# **Kruskal’s MST Algorithm**

## **Overview:**

Kruskal's algorithm finds the Minimum Spanning Tree (MST) of a graph, ensuring all vertices are connected with the minimum total edge weight.

## **Python Example:**

- **Kruskal’s MST Algorithm:**

```python
def find(parent, i):
    if parent[i] == i:
        return i
    else:
        return find(parent, parent[i])

def union(parent, rank, x, y):
    rootX = find(parent, x)
    rootY = find(parent, y)
    
    if rootX != rootY:
        if rank[rootX] > rank[rootY]:
            parent[rootY] = rootX
        elif rank[rootX] < rank[rootY]:
            parent[rootX] = rootY
        else:
            parent[rootY] = rootX
            rank[rootX] += 1

def kruskal(vertices, edges):
    parent = list(range(vertices))
    rank = [0] * vertices
    mst = []
    
    edges.sort(key=lambda x: x[2])
    
    for u, v, weight in edges:
        if find(parent, u) != find(parent, v):
            union(parent, rank, u, v)
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
print("Minimum Spanning Tree:", kruskal(vertices, edges))
```

## **Explanation:**
- **Kruskal’s Algorithm:** Sorts all edges by weight and adds the smallest edge to the MST while avoiding cycles using Union-Find data structure.
