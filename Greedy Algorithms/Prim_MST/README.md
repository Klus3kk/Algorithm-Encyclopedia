# **Prim’s MST Algorithm**

## **Overview:**

Prim’s Algorithm finds the Minimum Spanning Tree (MST) by starting from a node and growing the MST one edge at a time, always adding the smallest edge that connects a new node to the growing MST.

## **Python Example:**

```python
import heapq

def prim(vertices, edges):
    graph = {i: [] for i in range(vertices)}
    for u, v, weight in edges:
        graph[u].append((weight, v))
        graph[v].append((weight, u))

    min_heap = [(0, 0)]  # Start with vertex 0 and weight 0
    visited = set()
    mst_weight = 0
    mst_edges = []

    while min_heap:
        weight, u = heapq.heappop(min_heap)
        if u not in visited:
            visited.add(u)
            mst_weight += weight
            if weight > 0:
                mst_edges.append((weight, u))
            for edge_weight, v in graph[u]:
                if v not in visited:
                    heapq.heappush(min_heap, (edge_weight, v))

    return mst_edges, mst_weight

# Example usage
vertices = 4
edges = [
    (0, 1, 10),
    (0, 2, 6),
    (0, 3, 5),
    (1, 3, 15),
    (2, 3, 4)
]
mst_edges, mst_weight = prim(vertices, edges)
print("Prim's MST Edges:", mst_edges)
print("Total Weight of MST:", mst_weight)
```

## **Explanation:**
- **Graph Representation:** Converts edge list to adjacency list.
- **Priority Queue:** Keeps track of the minimum edge weights.
- **MST Construction:** Expands the MST by adding the smallest edge connecting to a new node.

