# **Hashing Techniques**

## **Overview**

Hashing is a technique used to map data to fixed-size values (hash codes) using a hash function. This technique is fundamental for implementing hashmaps, sets, and other data structures.

## **Where is it Used?**

- **Hashmaps:** Efficient key-value pair storage.
- **Sets:** Storing unique elements with fast lookups.
- **Caches:** Implementing hash-based caching mechanisms.

## **How Does it Work?**

A hash function takes an input (or key) and returns an integer (hash code). This hash code determines the index in a data structure where the key-value pair or element is stored.

## **Python Example**

Here’s an example of using Python's built-in `hash` function for hashing a string:

```python
def hash_string(s):
    return hash(s)

# Example usage
string = "example"
print("Hash code of '{}':".format(string), hash_string(string))
```

- **Explanation:**  
  The `hash` function computes a hash value for the given string, which can be used as a key in hash-based data structures.

