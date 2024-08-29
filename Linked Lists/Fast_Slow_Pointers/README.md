# **Fast and Slow Pointers**

## **Overview**

The fast and slow pointers technique is used for various algorithms on linked lists, such as detecting cycles or finding the middle of the list. This approach involves using two pointers that move at different speeds through the list.

## **Where is it Used?**

- **Cycle Detection:** Identifying if a linked list contains a cycle.
- **Finding Middle Element:** Locating the middle node of a linked list efficiently.
- **Linked List Operations:** Optimizing certain linked list operations.

## **How Does it Work?**

- **Cycle Detection:** Move one pointer (slow) by one step and the other pointer (fast) by two steps. If they meet, a cycle exists.
- **Finding Middle Element:** Move one pointer (slow) by one step and the other pointer (fast) by two steps. When the fast pointer reaches the end, the slow pointer will be at the middle.

## **Python Example**

Here's an example of cycle detection using fast and slow pointers:

```python
class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

def has_cycle(head):
    slow = head
    fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            return True
    return False

# Example usage
head = Node(1)
second = Node(2)
third = Node(3)
head.next = second
second.next = third
third.next = head  # Creates a cycle

print("Has cycle:", has_cycle(head))  # Output: True
```

- **Explanation:**  
  The `has_cycle` function uses two pointers moving at different speeds to detect if there's a cycle in the list.

