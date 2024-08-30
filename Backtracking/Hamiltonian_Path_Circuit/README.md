# **Hamiltonian Path and Circuit**

## **Overview:**

The Hamiltonian Path problem involves finding a path in a graph that visits each vertex exactly once. If the path is a cycle (returns to the starting vertex), it is called a Hamiltonian Circuit.

## **Python Example:**

```python
def hamiltonian_path(graph):
    def is_valid(v, pos, path):
        if graph[path[pos - 1]][v] == 0:
            return False
        if v in path:
            return False
        return True

    def solve(path, pos):
        if pos == len(graph):
            return True
        
        for v in range(len(graph)):
            if is_valid(v, pos, path):
                path[pos] = v
                if solve(path, pos + 1):
                    return True
                path[pos] = -1
        
        return False

    path = [-1] * len(graph)
    path[0] = 0
    if solve(path, 1):
        return path
    else:
        return None

# Example usage
graph = [[0, 1, 1, 1],
         [1, 0, 1, 1],
         [1, 1, 0, 1],
         [1, 1, 1, 0]]

path = hamiltonian_path(graph)
print("Hamiltonian Path:", path)
```

## **Explanation:**
- **Validity Check:** Ensure the next vertex is connected and not already in the path.
- **Backtracking:** Attempt to build the path vertex by vertex and backtrack if needed.

