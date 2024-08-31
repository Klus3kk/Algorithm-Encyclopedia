# **Networking Algorithms**

## **Overview:**

Networking Algorithms are used to efficiently manage and route data across networks. They optimize communication and ensure reliable data transfer.

## **Subtopics:**

- **Routing Algorithms:** Determine the best path for data to travel through a network.
- **Load Balancing:** Distributes workloads evenly across servers to optimize resource utilization.

## **Python Example (Dijkstra’s Shortest Path):**

```python
import heapq

def dijkstra(graph, start):
    distances = {node: float('inf') for node in graph}
    distances[start] = 0
    priority_queue = [(0, start)]
    
    while priority_queue:
        current_distance, u = heapq.heappop(priority_queue)
        
        if current_distance > distances[u]:
            continue
        
        for neighbor, weight in graph[u]:
            distance = current_distance + weight
            if distance < distances[neighbor]:
                distances[neighbor] = distance
                heapq.heappush(priority_queue, (distance, neighbor))
    
    return distances

# Example usage
graph = {
    'A': [('B', 1), ('C', 4)],
    'B': [('C', 2), ('D', 5)],
    'C': [('D', 1)],
    'D': []
}
print("Shortest paths from A:", dijkstra(graph, 'A'))
```

## **Explanation:**
- **Dijkstra’s Algorithm:** Finds the shortest path from a starting node to all other nodes in a weighted graph using a priority queue.

