# **Flood Fill Algorithm**

## **Overview:**

Flood fill algorithm is used to determine connected regions in a grid, such as in paint bucket tools in graphics applications.

## **Python Example:**

- **Flood Fill Algorithm:**

```python
def flood_fill(grid, start, new_color):
    rows, cols = len(grid), len(grid[0])
    original_color = grid[start[0]][start[1]]
    
    def fill(r, c):
        if (r < 0 or r >= rows or c < 0 or c >= cols or grid[r][c] != original_color):
            return
        grid[r][c] = new_color
        fill(r + 1, c)
        fill(r - 1, c)
        fill(r, c + 1)
        fill(r, c - 1)
    
    fill(start[0], start[1])
    return grid

# Example usage
grid = [
    [1, 1, 1],
    [1, 1, 0],
    [1, 0, 1]
]
start = (1, 1)
new_color = 2
print("Grid after flood fill:", flood_fill(grid, start, new_color))
```

## **Explanation:**
- **Flood Fill Algorithm:** A depth-first search approach to fill connected regions with a new color, used in image processing.

