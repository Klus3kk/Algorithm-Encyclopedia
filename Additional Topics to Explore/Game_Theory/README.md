# **Game Theory**

## **Overview:**

Game Theory studies mathematical models of strategic interaction among rational decision-makers. It includes algorithms for optimal decision-making in competitive scenarios.

## **Subtopics:**

- **Minimax Algorithm:** Used in decision-making for two-player games.
- **Alpha-Beta Pruning:** Optimizes the minimax algorithm by pruning branches.

## **Python Example (Minimax Algorithm):**

```python
def minimax(depth, node_index, maximizingPlayer, scores, h):
    if depth == h:
        return scores[node_index]
    if maximizingPlayer:
        return max(minimax(depth + 1, node_index * 2, False, scores, h),
                   minimax(depth + 1, node_index * 2 + 1, False, scores, h))
    else:
        return min(minimax(depth + 1, node_index * 2, True, scores, h),
                   minimax(depth + 1, node_index * 2 + 1, True, scores, h))

# Example usage
scores = [3, 5, 2, 9]
print("Optimal value:", minimax(0, 0, True, scores, 2))
```

## **Explanation:**
- **Minimax Algorithm:** Recursively evaluates all possible moves in a game, choosing the optimal move for the current player assuming the opponent plays optimally.

