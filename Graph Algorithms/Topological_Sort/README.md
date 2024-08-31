# **Topological Sort**

## **Overview:**

Topological sort orders the vertices of a directed acyclic graph (DAG) such that for every directed edge \( u \to v \), vertex \( u \) comes before \( v \).

## **Python Example:**

- **Topological Sort (DFS-Based):**

```python
from collections import defaultdict, deque

def topological_sort(graph):
    indegree = defaultdict(int)
    for u in graph:
        for v in graph[u]:
            indegree[v] += 1
    
    queue = deque([u for u in graph if indegree[u] == 0])
    topo_order = []
    
    while queue:
        u = queue.popleft()
        topo_order.append(u)
        
        for v in graph[u]:
            indegree[v] -= 1
            if indegree[v] == 0:
                queue.append(v)
    
    if len(topo_order) == len(graph):
        return topo_order
    else:
        raise ValueError("Graph is not a DAG")

# Example usage
graph = {
    5: [2, 0],
    4: [0, 1],
    2: [3],
    3: [1]
}
print("Topological sort:", topological_sort(graph))
```

## **Explanation:**
- **Topological Sort:** Uses Kahn’s Algorithm with an indegree array to maintain vertices with zero indegree, processing them to generate the sorted order.

