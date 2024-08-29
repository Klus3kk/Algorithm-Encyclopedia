# **Stack Operations**

## **Overview**

A stack is a linear data structure that follows the Last-In-First-Out (LIFO) principle. Elements can be added to or removed from only one end, known as the top of the stack.

## **Where is it Used?**

- **Function Call Management:** Tracking function calls and returns.
- **Expression Evaluation:** Evaluating mathematical expressions and parsing syntax.
- **Undo Mechanisms:** Implementing undo features in applications.

## **How Does it Work?**

- **Push:** Add an element to the top of the stack.
- **Pop:** Remove the top element from the stack.
- **Peek:** Retrieve the top element without removing it.

## **Python Example**

Here's an implementation of a stack using a Python list:

```python
class Stack:
    def __init__(self):
        self.stack = []

    def push(self, item):
        self.stack.append(item)

    def pop(self):
        if not self.is_empty():
            return self.stack.pop()
        raise IndexError("Pop from empty stack")

    def peek(self):
        if not self.is_empty():
            return self.stack[-1]
        raise IndexError("Peek from empty stack")

    def is_empty(self):
        return len(self.stack) == 0

# Example usage
stack = Stack()
stack.push(1)
stack.push(2)
print("Peek:", stack.peek())  # Output: 2
print("Pop:", stack.pop())    # Output: 2
print("Pop:", stack.pop())    # Output: 1
```

- **Explanation:**  
  The `Stack` class allows pushing and popping elements and checking if the stack is empty.

