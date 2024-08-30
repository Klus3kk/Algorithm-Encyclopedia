# **DP on Graphs**

## **Overview:**

DP on Graphs involves solving problems where the data structure is a graph, and solutions require optimizing over paths or subgraphs.

## **Python Example:**

- **Shortest Path in Weighted Graph (Bellman-Ford Algorithm):**

```python
def bellman_ford(n, edges, source):
    distance = [float('inf')] * n
    distance[source] = 0

    for _ in range(n - 1):
        for u, v, w in edges:
            if distance[u] != float('inf') and distance[u] + w < distance[v]:
                distance[v] = distance[u] + w

    for u, v, w in edges:
        if distance[u] != float('inf') and distance[u] + w < distance[v]:
            raise ValueError("Graph contains a negative weight cycle")

    return distance

# Example usage
n = 5
edges = [(0, 1, -1), (0, 2, 4), (1, 2, 3), (1, 3, 2), (1, 4, 2), (3, 2, 5), (3, 4, 1), (4, 3, -3)]
source = 0
print("Shortest paths from source:", bellman_ford(n, edges, source))
```

## **Explanation:**
- **Bellman-Ford Algorithm:** Computes the shortest path from a single source to all other vertices in a graph, handling negative weights and detecting negative weight cycles.

