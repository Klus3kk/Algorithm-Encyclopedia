# **Fractional Knapsack**

## **Overview:**

The Fractional Knapsack Problem allows items to be broken into smaller pieces to maximize the total value that fits into a knapsack of given capacity. It is solved using a greedy strategy by selecting items based on their value-to-weight ratio.

## **Python Example:**

```python
def fractional_knapsack(weights, values, capacity):
    items = sorted(zip(weights, values), key=lambda x: x[1]/x[0], reverse=True)
    
    total_value = 0
    for weight, value in items:
        if capacity == 0:
            break
        take_weight = min(weight, capacity)
        total_value += take_weight * (value / weight)
        capacity -= take_weight
    
    return total_value

# Example usage
weights = [10, 20, 30]
values = [60, 100, 120]
capacity = 50
print("Maximum value in Knapsack:", fractional_knapsack(weights, values, capacity))
```

## **Explanation:**
- **Value-to-Weight Ratio:** Items are sorted by their value-to-weight ratio.
- **Selection:** Items are added to the knapsack based on their ratio until the knapsack is full or all items are considered.
