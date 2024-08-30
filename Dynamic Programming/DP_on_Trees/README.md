# **DP on Trees**

## **Overview:**

DP on Trees involves solving problems where the data structure is a tree, and the solution requires combining results from subtrees.

## **Python Example:**

- **Tree Diameter:**

```python
def tree_diameter(n, edges):
    from collections import defaultdict, deque

    def bfs(start):
        visited = [-1] * n
        queue = deque([start])
        visited[start] = 0
        farthest_node = start
        while queue:
            node = queue.popleft()
            for neighbor, weight in graph[node]:
                if visited[neighbor] == -1:
                    visited[neighbor] = visited[node] + weight
                    queue.append(neighbor)
                    if visited[neighbor] > visited[farthest_node]:
                        farthest_node = neighbor
        return farthest_node, visited[farthest_node]

    graph = defaultdict(list)
    for u, v, w in edges:
        graph[u].append((v, w))
        graph[v].append((u, w))

    farthest_node, _ = bfs(0)
    _, diameter = bfs(farthest_node)
    
    return diameter

# Example usage
n = 5
edges = [(0, 1, 1), (1, 2, 2), (2, 3, 3), (3, 4, 4)]
print("Tree Diameter:", tree_diameter(n, edges))
```

## **Explanation:**
- **Tree Diameter:** Computes the longest path between any two nodes in a tree using two BFS traversals.

