# **Graph Representations**

## **Overview**

Graphs can be represented in various ways, including adjacency lists and adjacency matrices.

## **Where is it Used?**

- **Graph Algorithms:** Different representations affect the efficiency of algorithms.
- **Graph Storage:** Choosing an appropriate representation for data storage.

## **How Does it Work?**

- **Adjacency List:** List of neighbors for each vertex.
- **Adjacency Matrix:** 2D array where each entry indicates the presence of an edge between vertices.

## **Python Example**

Here’s an example of adjacency list and matrix representations:

```python
# Adjacency List
graph_list = {
    0: [1, 4],
    1: [0, 2, 4],
    2: [1, 3],
    3: [2, 4],
    4: [0, 1, 3]
}

# Adjacency Matrix
graph_matrix = [
    [0, 1, 0, 0, 1],
    [1, 0, 1, 0, 1],
    [0, 1, 0, 1, 0],
    [0, 0

, 1, 0, 1],
    [1, 1, 0, 1, 0]
]

print("Adjacency List:", graph_list)
print("Adjacency Matrix:", graph_matrix)
```

- **Explanation:**  
  Adjacency list and matrix represent the same graph in different formats.

