# **Binary Search Tree (BST) Operations**

## **Overview**

A Binary Search Tree (BST) is a binary tree where each node has at most two children, and the left child is less than the node, while the right child is greater.

## **Where is it Used?**

- **Searching:** Efficiently finding elements in a sorted dataset.
- **Insertion and Deletion:** Managing dynamic sets of data.
- **Range Queries:** Finding elements within a specified range.

## **How Does it Work?**

- **Insertion:** Insert nodes while maintaining the BST property.
- **Deletion:** Remove nodes and reorganize the tree to maintain the BST property.
- **Search:** Traverse the tree to find a specific value.

## **Python Example**

Here’s an example of BST operations:

```python
class BSTNode:
    def __init__(self, key):
        self.left = None
        self.right = None
        self.value = key

def insert(root, key):
    if root is None:
        return BSTNode(key)
    if key < root.value:
        root.left = insert(root.left, key)
    else:
        root.right = insert(root.right, key)
    return root

def search(root, key):
    if root is None or root.value == key:
        return root
    if key < root.value:
        return search(root.left, key)
    return search(root.right, key)

# Example usage
root = BSTNode(10)
root = insert(root, 5)
root = insert(root, 15)
print("Search 5:", search(root, 5) is not None)  # Output: True
print("Search 20:", search(root, 20) is not None) # Output: False
```

- **Explanation:**  
  The `insert` function adds elements to the BST, while the `search` function locates a specific value.

