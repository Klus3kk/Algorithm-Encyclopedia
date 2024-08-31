# **Bellman-Ford Algorithm**

## **Overview:**

Bellman-Ford algorithm finds the shortest path from a single source vertex to all other vertices in a graph, even with negative weights, but not negative weight cycles.

## **Python Example:**

- **Bellman-Ford Algorithm:**

```python
def bellman_ford(vertices, edges, start):
    distances = [float('inf')] * vertices
    distances[start] = 0
    
    for _ in range(vertices - 1):
        for u, v, weight in edges:
            if distances[u] != float('inf') and distances[u] + weight < distances[v]:
                distances[v] = distances[u] + weight
    
    # Check for negative weight cycles
    for u, v, weight in edges:
        if distances[u] != float('inf') and distances[u] + weight < distances[v]:
            raise ValueError("Graph contains negative weight cycle")
    
    return distances

# Example usage
vertices = 5
edges = [
    (0, 1, -1),
    (0, 2, 4),
    (1, 2, 3),
    (1, 3, 2),
    (1, 4, 2),
    (3, 2, 5),
    (3, 4, 1),
    (4, 3, -3)
]
print("Shortest paths from vertex 0:", bellman_ford(vertices, edges, 0))
```

## **Explanation:**
- **Bellman-Ford Algorithm:** Relaxes edges up to \( V-1 \) times to find shortest paths and then checks for negative weight cycles.

