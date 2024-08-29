# **Reversing a Linked List**

## **Overview**

Reversing a linked list involves reversing the order of its nodes. This operation can be done iteratively or recursively.

## **Where is it Used?**

- **Data Manipulation:** Reversing the order of elements for various algorithms.
- **Linked List Operations:** Performing reverse operations as part of linked list management.

## **How Does it Work?**

- **Iterative Approach:** Traverse the list while reversing the `next` pointers of each node.
- **Recursive Approach:** Recursively reverse the list and adjust pointers.

## **Python Example**

Here’s an iterative approach to reversing a singly linked list:

```python
class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

def reverse_linked_list(head):
    prev = None
    current = head
    while current:
        next_node = current.next
        current.next = prev
        prev = current
        current = next_node
    return prev

# Example usage
head = Node(1)
second = Node(2)
third = Node(3)
head.next = second
second.next = third

reversed_head = reverse_linked_list(head)
current = reversed_head
reversed_list = []
while current:
    reversed_list.append(current.value)
    current = current.next

print("Reversed Linked List:", reversed_list)  # Output: [3, 2, 1]
```

- **Explanation:**  
  The `reverse_linked_list` function iterates through the list, reversing the links between nodes.

