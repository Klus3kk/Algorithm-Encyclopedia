# **Graph Traversal**

## **Overview**

Graph traversal algorithms visit all nodes in a graph to explore its structure. Common methods include Depth-First Search (DFS) and Breadth-First Search (BFS).

## **Where is it Used?**

- **Searching and Pathfinding:** Finding paths and components in graphs.
- **Network Analysis:** Analyzing connections and clusters.
- **Web Crawlers:** Traversing web pages and links.

## **How Does it Work?**

- **Depth-First Search (DFS):** Explore as far as possible along each branch before backtracking.
- **Breadth-First Search (BFS):** Explore all nodes at the present depth level before moving on to nodes at the next depth level.

## **Python Example**

Here’s an example of BFS and DFS in Python:

```python
from collections import deque

def bfs(graph, start):
    visited = set()
    queue = deque([start])
    while queue:
        vertex = queue.popleft()
        if vertex not in visited:
            visited.add(vertex)
            queue.extend(neighbor for neighbor in graph[vertex] if neighbor not in visited)
    return visited

def dfs(graph, start, visited=None):
    if visited is None:
        visited = set()
    visited.add(start)
    for neighbor in graph[start]:
        if neighbor not in visited:
            dfs(graph, neighbor, visited)
    return visited

# Example usage
graph = {
    1: [2, 3],
    2: [1, 4, 5],
    3: [1, 6],
    4: [2],
    5: [2],
    6: [3]
}
print("BFS:", bfs(graph, 1))  # Output: {1, 2, 3, 4, 5, 6}
print("DFS:", dfs(graph, 1))  # Output: {1, 2, 3, 4, 5, 6}
```

- **Explanation:**  
  The `bfs` function uses a queue for breadth-first traversal, while the `dfs` function uses recursion for depth-first traversal.

