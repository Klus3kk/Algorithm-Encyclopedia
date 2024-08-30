# **Strongly Connected Components**

## **Overview**

Strongly Connected Components (SCCs) are subgraphs where every vertex is reachable from every other vertex within the subgraph.

## **Where is it Used?**

- **Network Analysis:** Identifying tightly knit groups in networks.
- **Dependency Resolution:** Managing dependencies in complex systems.

## **How Does it Work?**

- **Kosaraju's Algorithm:** Two-pass DFS to identify SCCs.
- **Tarjan's Algorithm:** DFS-based approach with low-link values to find SCCs.

## **Python Example**

Here’s an example of Kosaraju’s Algorithm:

```python
from collections import defaultdict, deque

class GraphSCC:
    def __init__(self, vertices):
        self.graph = defaultdict(list)
        self.V = vertices

    def add_edge(self, u, v):
        self.graph[u].append(v)

    def dfs(self, v, visited, stack):
        visited[v] = True
        for neighbor in self.graph[v]:
            if not visited[neighbor]:
                self.dfs(neighbor, visited, stack)
        stack.append(v)

    def transpose(self):
        g_t = GraphSCC(self.V)
        for v in self.graph:
            for neighbor in self.graph[v]:
                g_t.add_edge(neighbor, v)
        return g_t

    def kosaraju(self):
        stack = []
        visited = [False] * self.V
        for i in range(self.V):
            if not visited[i]:
                self.dfs(i, visited, stack)
        g_t = self.transpose()
        visited = [False] * self.V
        sccs = []
        while stack:
            v = stack.pop()
            if not visited[v]:
                scc = []
                g_t.dfs(v, visited, scc)
                sccs.append(scc)
        return sccs

# Example usage
g = GraphSCC(5)
g.add_edge(0, 2)
g.add_edge(2, 1)
g.add_edge(1, 0)
g.add_edge(1, 3)
g.add_edge(3, 4)
print("Strongly Connected Components:", g.kosaraju())  # Output: Strongly Connected Components: [[0, 1, 2], [3], [4]]
```

- **Explanation:**  
  Kosaraju’s Algorithm uses two DFS passes to find SCCs in a graph.

