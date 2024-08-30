# **Subset Sum**

## **Overview:**

The Subset Sum problem involves finding a subset of numbers that add up to a target sum. It is used in various applications, including resource allocation and decision making.

## **Python Example:**

```python
def subset_sum(arr, target):
    def backtrack(start, path, total):
        if total == target:
            solutions.append(path)
            return
        if total > target or start >= len(arr):
            return
        
        for i in range(start, len(arr)):
            backtrack(i + 1, path + [arr[i]], total + arr[i])
    
    solutions = []
    backtrack(0, [], 0)
    return solutions

# Example usage
arr = [3, 34, 4, 12, 5, 2]
target = 9
print("Subsets with sum", target, ":", subset_sum(arr, target))
```

## **Explanation:**
- **Recursive Search:** Try including each number in the subset and track if the sum reaches the target.
- **Backtracking:** Continue until all subsets are explored or the target sum is achieved.

