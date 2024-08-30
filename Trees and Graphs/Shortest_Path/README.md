# **Shortest Path**

## **Overview**

Shortest Path algorithms find the shortest path between nodes in a graph. Common algorithms include Dijkstra’s and Bellman-Ford.

## **Where is it Used?**

- **Navigation Systems:** Finding the shortest route between locations.
- **Network Routing:** Efficiently directing data in networks.

## **How Does it Work?**

- **Dijkstra’s Algorithm:** Use a priority queue to find the shortest path from a source node to all other nodes.
- **Bellman-Ford Algorithm:** Handle graphs with negative weights and detect negative weight cycles.

## **Python Example**

Here’s an example of Dijkstra’s Algorithm:

```python
import heapq

def dijkstra(graph, start):
    heap = [(0, start)]
    distances = {vertex: float('infinity') for vertex in graph}
    distances[start] = 0
    while heap:
        current_distance, current_vertex = heapq.heappop(heap)
        if current_distance > distances[current_vertex]:
            continue
        for neighbor, weight in graph[current_vertex]:
            distance = current_distance + weight
            if distance < distances[neighbor]:
                distances[neighbor] = distance
                heapq.heappush(heap, (distance, neighbor))
    return distances

# Example usage
graph = {
    'A': [('B', 1), ('C', 4)],
    'B': [('A', 1), ('C', 2), ('D', 5)],
    'C': [('A', 4), ('B', 2), ('D', 1)],
    'D': [('B', 5), ('C', 1)]
}
print("Shortest paths from A:", dijkstra(graph, 'A'))
# Output: Shortest paths from A: {'A': 0, 'B': 1, 'C': 3, 'D': 4}
```

- **Explanation:**  
  Dijkstra’s Algorithm uses a priority queue to efficiently find shortest paths in a graph.

