# **Graph Coloring**

## **Overview:**

Graph Coloring is the problem of assigning colors to vertices of a graph such that no two adjacent vertices have the same color. It is used in scheduling, map coloring, and resource allocation.

## **Python Example:**

```python
def graph_coloring(graph, m):
    def is_valid(v, color, c):
        for i in range(len(graph)):
            if graph[v][i] == 1 and color[i] == c:
                return False
        return True

    def solve(m, color, v):
        if v == len(graph):
            return True
        
        for c in range(1, m + 1):
            if is_valid(v, color, c):
                color[v] = c
                if solve(m, color, v + 1):
                    return True
                color[v] = 0
        
        return False

    color = [0] * len(graph)
    if solve(m, color, 0):
        return color
    else:
        return None

# Example usage
graph = [[0, 1, 1, 1],
         [1, 0, 1, 0],
         [1, 1, 0, 1],
         [1, 0, 1, 0]]

m = 3  # Number of colors
colors = graph_coloring(graph, m)
print("Graph Coloring:", colors)
```

## **Explanation:**
- **Validity Check:** Ensure no two adjacent vertices have the same color.
- **Backtracking:** Try coloring each vertex and backtrack if no valid coloring is possible.
