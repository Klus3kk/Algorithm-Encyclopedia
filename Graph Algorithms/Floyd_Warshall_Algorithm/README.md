# **Floyd-Warshall Algorithm**

## **Overview:**

Floyd-Warshall algorithm finds shortest paths between all pairs of vertices in a graph, handling negative weights but not negative weight cycles.

## **Python Example:**

- **Floyd-Warshall Algorithm:**

```python
def floyd_warshall(graph):
    num_vertices = len(graph)
    dist = [[float('inf')] * num_vertices for _ in range(num_vertices)]
    
    for u in range(num_vertices):
        for v in range(num_vertices):
            if u == v:
                dist[u][v] = 0
            elif (u, v) in graph:
                dist[u][v] = graph[(u, v)]
    
    for k in range(num_vertices):
        for i in range(num_vertices):
            for j in range(num_vertices):
                if dist[i][k] + dist[k][j] < dist[i][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]
    
    return dist

# Example usage
graph = {
    (0, 1): 3,
    (1, 2): 1,
    (2, 0): 2,
    (1, 3): 4,
    (3, 2): 6
}
print("Shortest paths matrix:", floyd_warshall(graph))
```

## **Explanation:**
- **Floyd-Warshall Algorithm:** Uses dynamic programming to update shortest paths between all pairs of vertices by considering each vertex as an intermediate point.

