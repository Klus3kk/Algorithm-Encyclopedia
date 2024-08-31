# **Bipartite Graph Checking**

## **Overview:**

A bipartite graph can be colored using two colors such that no two adjacent vertices share the same color. This check can be performed using BFS or DFS.

## **Python Example:**

- **Bipartite Graph Check Using BFS:**

```python
from collections import deque

def is_bipartite(graph):
    color = {}
    for node in graph:
        color[node] = None
    
    def bfs(start):
        queue = deque([start])
        color[start] = 0
        
        while queue:
            u = queue.popleft()
            for v in graph[u]:
                if color[v] is None:
                    color[v] = 1 - color[u]
                    queue.append(v)
                elif color[v] == color[u]:
                    return False
        return True
    
    for node in graph:
        if color[node] is None:
            if not bfs(node):
                return False
    
    return True

# Example usage
graph = {
    0: [1, 3],
    1: [0, 2],
    2: [1, 3],
    3: [0, 2]
}
print("Is the graph bipartite?", is_bipartite(graph))
```

## **Explanation:**
- **Bipartite Graph Checking:** Uses BFS to try and color the graph with two colors, checking for adjacent vertices with the same color.

