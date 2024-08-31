# **Strongly Connected Components**

## **Overview:**

Strongly Connected Components (SCCs) are subgraphs where every vertex is reachable from every other vertex within the subgraph. Kosaraju’s and Tarjan’s algorithms are commonly used to find SCCs.

## **Python Example:**

- **Finding SCCs Using Kosaraju’s Algorithm:**

```python
def kosaraju_scc(graph):
    def dfs(v, visited, stack):
        visited[v] = True
        for neighbor in graph[v]:
            if not visited[neighbor]:
                dfs(neighbor, visited, stack)
        stack.append(v)
    
    def reverse_dfs(v, visited, scc):
        visited[v] = True
        scc.append(v)
        for neighbor in reversed_graph[v]:
            if not visited[neighbor]:
                reverse_dfs(neighbor, visited, scc)
    
    stack = []
    visited = {v: False for v in graph}
    
    for v in graph:
        if not visited[v]:
            dfs(v, visited, stack)
    
    reversed_graph = {v: [] for v in graph}
    for v in graph:
        for neighbor in graph[v]:
            reversed_graph[neighbor].append(v)
    
    visited = {v: False for v in graph}
    sccs = []
    
    while stack:
        v = stack.pop()
        if not visited[v]:
            scc = []
            reverse_dfs(v, visited, scc)
            sccs.append(scc)
    
    return sccs

# Example usage
graph = {
    0: [1],
    1: [2],
    2: [0, 3],
    3: [4],
    4: [5],
    5: [3]
}
print("Strongly Connected Components:", kosaraju_scc(graph))
```

## **Explanation:**
- **Kosaraju’s Algorithm:** Uses two DFS passes—one on the original graph to order vertices by finish time and another on the reversed graph to find SCCs.

