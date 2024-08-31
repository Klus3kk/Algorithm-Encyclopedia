# **Lee Algorithm**

## **Overview:**

Lee algorithm is used for finding the shortest path in an unweighted grid or maze by performing a breadth-first search.

## **Python Example:**

- **Lee Algorithm (BFS for Grid):**

```python
from collections import deque

def lee_algorithm(grid, start, goal):
    rows, cols = len(grid), len(grid[0])
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    queue = deque([start])
    distance = [[float('inf')] * cols for _ in range(rows)]
    distance[start[0]][start[1]] = 0
    
    while queue:
        r, c = queue.popleft()
        
        if (r, c) == goal:
            return distance[r][c]
        
        for dr, dc in directions:
            nr, nc = r + dr, c + dc
            if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == 0 and distance[nr][nc] == float('inf'):
                distance[nr][nc] = distance[r][c] + 1
                queue.append((nr, nc))
    
    return -1  # Goal not reachable

# Example usage
grid = [
    [0, 1, 0, 0],
    [0, 1, 0, 0],
    [0, 0, 0, 0],
    [0, 1, 1, 0]
]
start = (0, 0)
goal = (3, 3)
print("Shortest path length:", lee_algorithm(grid, start, goal))
```

## **Explanation:**
- **Lee Algorithm:** A BFS approach to find the shortest path in an unweighted grid, considering movement in four directions.

