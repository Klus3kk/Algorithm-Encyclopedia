# **Network Flow**

## **Overview:**

Network flow algorithms handle problems where flow must be sent through a network with capacities on edges, such as the Max-Flow problem.

## **Python Example:**

- **Ford-Fulkerson Algorithm for Max-Flow:**

```python
from collections import defaultdict

class MaxFlow:
    def __init__(self, vertices):
        self.graph = defaultdict(dict)
        self.vertices = vertices
    
    def add_edge(self, u, v, capacity):
        self.graph[u][v] = capacity
        self.graph[v][u] = 0
    
    def bfs(self, source, sink, parent):
        visited = set()
        queue = [source]
        visited.add(source)
        
        while queue:
            u = queue.pop(0)
            if u == sink:
                return True
            for v, cap in self.graph[u].items():
                if v not in visited and cap > 0:
                    parent[v] = u
                    visited.add(v)
                    queue.append(v)
                    if v == sink:
                        return True
        return False
    
    def ford_fulkerson(self, source, sink):
        parent = {}
        max_flow = 0
        
        while self.bfs(source, sink, parent):
            path_flow = float('inf')
            s = sink
            
            while s != source:
                path_flow = min(path_flow, self.graph[parent[s]][s])
                s = parent[s]
            
            v = sink
            while v != source:
                u = parent[v]
                self.graph[u][v] -= path_flow
                self.graph[v][u] += path_flow
                v = parent[v]
            
            max_flow += path_flow
        
        return max_flow

# Example usage
mf = MaxFlow(6)
mf.add_edge(0, 1, 16)
mf.add_edge(0, 2, 13)
mf.add_edge(1, 2, 10)
mf.add_edge(1, 3, 12)
mf.add_edge(2, 1, 4)
mf.add_edge(2, 4, 14)
mf.add_edge(3, 2, 9)
mf.add_edge(3, 5, 20)
mf.add_edge(4, 3, 7)
mf.add_edge(4, 5, 4)

print("Max Flow:", mf.ford_fulkerson(0, 5))
```

## **Explanation:**
- **Ford-Fulkerson Algorithm:** Uses BFS to find augmenting paths and adjusts capacities to compute the maximum flow from source to sink.

