# **Articulation Points and Bridges**

## **Overview:**

Articulation points (or cut vertices) are nodes in a graph whose removal increases the number of connected components. Bridges are edges whose removal increases the number of connected components.

## **Python Example:**

- **Finding Articulation Points and Bridges (DFS-Based):**

```python
def find_articulation_points_and_bridges(graph):
    def dfs(u, parent):
        nonlocal time
        children = 0
        visited[u] = True
        discovery[u] = low[u] = time
        time += 1
        
        for v in graph[u]:
            if not visited[v]:
                children += 1
                parent[v] = u
                dfs(v, parent)
                
                low[u] = min(low[u], low[v])
                
                if parent[u] is None and children > 1:
                    articulation_points.add(u)
                if parent[u] is not None and low[v] >= discovery[u]:
                    articulation_points.add(u)
                if low[v] > discovery[u]:
                    bridges.append((u, v))
            elif v != parent[u]:
                low[u] = min(low[u], discovery[v])
    
    time = 0
    visited = {u: False for u in graph}
    discovery = {u: float('inf') for u in graph}
    low = {u: float('inf') for u in graph}
    parent = {u: None for u in graph}
    articulation_points = set()
    bridges = []
    
    for u in graph:
        if not visited[u]:
            dfs(u, parent)
    
    return articulation_points, bridges

# Example usage
graph = {
    0: [1, 2],
    1: [0, 2, 3],
    2: [0, 1, 3],
    3: [1, 2, 4],
    4: [3]
}
print("Articulation points:", find_articulation_points_and_bridges(graph)[0])
print("Bridges:", find_articulation_points_and_bridges(graph)[1])
```

## **Explanation:**
- **Articulation Points and Bridges:** Uses DFS to identify articulation points and bridges by comparing discovery and low values.

