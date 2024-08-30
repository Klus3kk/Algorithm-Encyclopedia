# **Topological Sort**

## **Overview**

Topological Sort orders vertices in a Directed Acyclic Graph (DAG) such that for every directed edge `uv`, vertex `u` comes before vertex `v`.

## **Where is it Used?**

- **Task Scheduling:** Ordering tasks with dependencies.
- **Course Prerequisites:** Determining course order based on prerequisites.

## **How Does it Work?**

- **Kahn’s Algorithm:** Uses in-degrees to build the topological order.
- **DFS-Based Algorithm:** Use DFS to build the topological order by processing nodes after exploring all their descendants.

## **Python Example**

Here’s an example of Topological Sort using DFS:

```python
from collections import defaultdict, deque

def topological_sort(graph):
    in_degree = {u: 0 for u in graph}
    for u in graph:
        for v in graph[u]:
            in_degree[v] += 1

    queue = deque([u for u in in_degree if in_degree[u] == 0])
    result = []
    
    while queue:
        u = queue.popleft()
        result.append(u)
        for v in graph[u]:
            in_degree[v] -= 1
            if in_degree[v] == 0:
                queue.append(v)
    
    return result

# Example usage
graph = {
    'A': ['B', 'C'],
    'B': ['D'],
    'C': ['D'],
    'D': []
}
print("Topological Sort:", topological_sort(graph))  # Output: Topological Sort: ['A', 'B', 'C', 'D']
```

- **Explanation:**  
  The `topological_sort` function finds an ordering of vertices in a DAG.

