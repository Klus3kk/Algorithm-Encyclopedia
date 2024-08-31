# **Greedy Techniques**

## **Overview:**

Greedy Techniques involve making locally optimal choices at each step with the hope of finding a global optimum. These algorithms are often used for optimization problems where a solution is built incrementally.

## **Subtopics:**

- **Activity Selection:** Selects the maximum number of non-overlapping activities.
- **Fractional Knapsack:** Maximizes the total value in a knapsack with fractional item inclusion allowed.
- **Huffman Coding:** Constructs an optimal prefix-free code for data compression.

## **Python Example (Activity Selection):**

```python
def activity_selection(start, end):
    n = len(start)
    activities = sorted(range(n), key=lambda i: end[i])
    result = [activities[0]]

    for i in range(1, n):
        if start[activities[i]] >= end[result[-1]]:
            result.append(activities[i])

    return result

# Example usage
start = [1, 3, 0, 5, 8, 5]
end = [2, 4, 6, 7, 9, 9]
selected_activities = activity_selection(start, end)
print("Selected activities:", selected_activities)
```

## **Explanation:**
- **Activity Selection:** Uses a greedy approach to select the maximum number of non-overlapping activities based on their finish times.

## **Python Example (Fractional Knapsack):**

```python
def fractional_knapsack(weights, values, capacity):
    items = [(values[i] / weights[i], weights[i], values[i]) for i in range(len(weights))]
    items.sort(reverse=True)

    total_value = 0
    for value_per_weight, weight, value in items:
        if capacity >= weight:
            capacity -= weight
            total_value += value
        else:
            total_value += value_per_weight * capacity
            break

    return total_value

# Example usage
weights = [10, 20, 30]
values = [60, 100,

 120]
capacity = 50
print("Maximum value in knapsack:", fractional_knapsack(weights, values, capacity))
```

## **Explanation:**
- **Fractional Knapsack:** Uses a greedy strategy to maximize the total value by allowing fractional item inclusion.

## **Python Example (Huffman Coding):**

```python
import heapq
from collections import Counter, defaultdict

class Node:
    def __init__(self, char, freq):
        self.char = char
        self.freq = freq
        self.left = None
        self.right = None

    def __lt__(self, other):
        return self.freq < other.freq

def huffman_encoding(data):
    frequency = Counter(data)
    priority_queue = [Node(char, freq) for char, freq in frequency.items()]
    heapq.heapify(priority_queue)

    while len(priority_queue) > 1:
        left = heapq.heappop(priority_queue)
        right = heapq.heappop(priority_queue)
        merged = Node(None, left.freq + right.freq)
        merged.left = left
        merged.right = right
        heapq.heappush(priority_queue, merged)

    def build_code_tree(node, prefix='', codebook=defaultdict()):
        if node:
            if node.char is not None:
                codebook[node.char] = prefix
            build_code_tree(node.left, prefix + '0', codebook)
            build_code_tree(node.right, prefix + '1', codebook)
        return codebook

    root = priority_queue[0]
    huffman_code = build_code_tree(root)

    encoded_data = ''.join(huffman_code[char] for char in data)
    return encoded_data, huffman_code

# Example usage
data = "this is an example for huffman encoding"
encoded_data, huffman_code = huffman_encoding(data)
print("Encoded data:", encoded_data)
print("Huffman code:", huffman_code)
```

## **Explanation:**
- **Huffman Coding:** Uses a greedy algorithm to build an optimal prefix-free code for data compression based on character frequencies.

