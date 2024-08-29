# **Merging Two Linked Lists**

## **Overview**

Merging two linked lists involves combining them into a single sorted linked list. This is often used in merge sort algorithms or combining two sorted lists.

## **Where is it Used?**

- **Merge Sort:** Combining sorted sublists.
- **Data Merging:** Merging data from multiple sources into a single list.

## **How Does it Work?**

Traverse both linked lists and merge them into a new sorted list by comparing nodes.

## **Python Example**

Here’s how to merge two sorted singly linked lists:

```python
class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

def merge_sorted_lists(list1, list2):
    dummy = Node(0)
    tail = dummy
    while list1 and list2:
        if list1.value < list2.value:
            tail.next = list1
            list1 = list1.next
        else:
            tail.next = list2
            list2 = list2.next
        tail = tail.next
    tail.next = list1 if list1 else list2
    return dummy.next

# Example usage
list1 = Node(1)
list1.next = Node(3)
list1.next.next = Node(5)

list2 = Node(2)
list2.next = Node(4)
list2.next.next = Node(6)

merged_list = merge_sorted_lists(list1, list2)
current = merged_list
merged_values = []
while current:
    merged_values.append(current.value)
    current = current.next

print("Merged Linked List:", merged_values)  # Output: [1, 2, 3, 4, 5, 6]
```

- **Explanation:**  
  The `merge_sorted_lists` function merges two sorted lists into a new sorted list.

