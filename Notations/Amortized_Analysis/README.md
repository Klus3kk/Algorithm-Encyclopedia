# **Amortized Analysis**

## **Overview**

Amortized analysis looks at the average time per operation over a sequence of operations, ensuring that occasional expensive operations are balanced by many cheap ones.

## **Where is it Used?**

- **Dynamic Arrays:** To analyze the time complexity of operations like append.
- **Data Structures:** In structures like hash tables or dynamic arrays where operations can have varying costs.

## **How Does it Work?**

The idea is to spread the cost of expensive operations across multiple cheaper ones. For instance:
- Appending to a dynamic array (which occasionally needs resizing) typically has an amortized O(1) complexity.

## **Python Example**

Consider appending elements to a list:

```python
arr = []
for i in range(1000):
    arr.append(i)

print(len(arr))
```

- **Complexity Analysis:**  
  While resizing the list occasionally requires O(n) time, the amortized time complexity of appending elements remains O(1).
