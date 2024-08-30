# **Articulation Points and Bridges**

## **Overview**

Articulation points and bridges are critical components whose removal would increase the number of connected components in a graph.

## **Where is it Used?**

- **Network Reliability:** Identifying critical nodes and edges in networks.
- **Graph Connectivity:** Analyzing and maintaining graph connectivity.

## **How Does it Work?**

- **Articulation Points:** Nodes whose removal increases the number of connected components.
- **Bridges:** Edges whose removal increases the number of connected components.

## **Python Example**

Here’s an example of finding articulation points using DFS:

```python
from collections import defaultdict

class Graph:
    def __init__(self, vertices):
        self.graph = defaultdict(list)
        self.V = vertices
        self.time = 0

    def add_edge(self, u, v):
        self.graph[u].append(v)
        self.graph[v].append(u)

    def articulation_points_util(self, u, visited, parent, low, disc, ap):
        children = 0
        visited[u] = True
        disc[u] = low[u] = self.time
        self.time += 1

        for v in self.graph[u]:
            if not visited[v]:
                parent[v] = u
                children += 1
                self.articulation_points_util(v, visited, parent, low, disc, ap)
                low[u] = min(low[u], low[v])
                if parent[u] is None and children > 1:
                    ap.add(u)
                if parent[u] is not None and low[v] >= disc[u]:
                    ap.add(u)
            elif v != parent[u]:
                low[u] = min(low[u], disc[v])

    def articulation_points(self):
        visited = [False] * self.V
        disc = [float('inf')] * self.V
        low = [float('inf')] * self.V
        parent = [None] * self.V
        ap = set()
        for i in range(self.V):
            if not visited[i]:
                self.articulation_points_util(i, visited, parent, low, disc, ap)
        return ap

# Example usage
g = Graph(5)
g.add_edge(0, 1)
g.add_edge(1, 2)
g.add_edge(2, 0)
g.add_edge(1, 3)
g.add_edge(3, 4)
print("Articulation Points:", g.articulation_points())  # Output: Articulation Points: {1}
```

- **Explanation:**  
  The `articulation_points` function finds nodes critical to maintaining graph connectivity.

