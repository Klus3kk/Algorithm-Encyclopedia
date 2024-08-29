# **Counting Frequencies**

## **Overview**

Counting frequencies involves determining how often each element appears in a collection. This is commonly done using hashmaps or dictionaries.

## **Where is it Used?**

- **Data Analysis:** Analyzing occurrences of items in datasets.
- **Text Processing:** Counting word frequencies in documents.

## **How Does it Work?**

Each element is used as a key in a hashmap, and its count is updated as elements are encountered.

## **Python Example**

Here’s an example of counting frequencies using a dictionary:

```python
def count_frequencies(arr):
    frequency_map = {}
    for item in arr:
        if item in frequency_map:
            frequency_map[item] += 1
        else:
            frequency_map[item] = 1
    return frequency_map

# Example usage
arr = ["apple", "banana", "apple", "orange", "banana", "banana"]
print("Frequencies:", count_frequencies(arr))  # Output: {'apple': 2, 'banana': 3, 'orange': 1}
```

- **Explanation:**  
  The `count_frequencies` function creates a frequency map where each unique element in the array is mapped to its count.

