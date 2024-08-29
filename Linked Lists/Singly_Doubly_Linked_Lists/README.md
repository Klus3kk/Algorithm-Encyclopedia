# **Singly and Doubly Linked Lists**

## **Overview**

A linked list is a collection of nodes where each node contains a value and a reference (or link) to the next node in the sequence. Singly linked lists have nodes that point only to the next node, while doubly linked lists have nodes that point to both the next and the previous nodes.

## **Where is it Used?**

- **Dynamic Data Storage:** Efficient insertion and deletion of elements.
- **Data Structures:** Building blocks for other data structures like stacks and queues.
- **Navigational Systems:** Implementing data structures requiring bidirectional traversal.

## **How Does it Work?**

- **Singly Linked List:** Each node has a `value` and a `next` pointer.
- **Doubly Linked List:** Each node has a `value`, `next` pointer, and a `prev` pointer.

## **Python Example**

Here’s an example implementation of a singly linked list:

```python
class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

class SinglyLinkedList:
    def __init__(self):
        self.head = None

    def append(self, value):
        new_node = Node(value)
        if not self.head:
            self.head = new_node
            return
        current = self.head
        while current.next:
            current = current.next
        current.next = new_node

    def display(self):
        values = []
        current = self.head
        while current:
            values.append(current.value)
            current = current.next
        return values

# Example usage
linked_list = SinglyLinkedList()
linked_list.append(1)
linked_list.append(2)
linked_list.append(3)
print("Singly Linked List:", linked_list.display())  # Output: [1, 2, 3]
```

- **Explanation:**  
  The `SinglyLinkedList` class allows appending values and displaying the list. Each node is linked to the next, forming a chain.

