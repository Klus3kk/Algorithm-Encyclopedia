# **Dijkstra’s Shortest Path**

## **Overview:**

Dijkstra's algorithm finds the shortest path from a source node to all other nodes in a weighted graph with non-negative weights.

## **Python Example:**

- **Dijkstra’s Algorithm:**

```python
import heapq

def dijkstra(graph, start):
    distances = {vertex: float('infinity') for vertex in graph}
    distances[start] = 0
    priority_queue = [(0, start)]
    
    while priority_queue:
        current_distance, current_vertex = heapq.heappop(priority_queue)
        
        if current_distance > distances[current_vertex]:
            continue
        
        for neighbor, weight in graph[current_vertex].items():
            distance = current_distance + weight
            if distance < distances[neighbor]:
                distances[neighbor] = distance
                heapq.heappush(priority_queue, (distance, neighbor))
    
    return distances

# Example usage
graph = {
    'A': {'B': 1, 'C': 4},
    'B': {'A': 1, 'C': 2, 'D': 5},
    'C': {'A': 4, 'B': 2, 'D': 1},
    'D': {'B': 5, 'C': 1}
}
print("Shortest paths from 'A':", dijkstra(graph, 'A'))
```

## **Explanation:**
- **Dijkstra’s Algorithm:** Uses a priority queue to explore the shortest known distance to each node. Updates distances to neighboring nodes if a shorter path is found.

