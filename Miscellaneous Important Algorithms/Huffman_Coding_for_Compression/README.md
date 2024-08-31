# **Huffman Coding for Compression**

## **Overview:**

Huffman Coding is a compression algorithm that assigns variable-length codes to input characters based on their frequencies. More frequent characters get shorter codes.

## **Python Example:**

```python
from heapq import heapify, heappush, heappop
from collections import defaultdict, Counter

class Node:
    def __init__(self, char, freq):
        self.char = char
        self.freq = freq
        self.left = None
        self.right = None

    def __lt__(self, other):
        return self.freq < other.freq

def huffman_coding(freqs):
    heap = [Node(char, freq) for char, freq in freqs.items()]
    heapify(heap)
    
    while len(heap) > 1:
        left = heappop(heap)
        right = heappop(heap)
        merged = Node(None, left.freq + right.freq)
        merged.left = left
        merged.right = right
        heappush(heap, merged)
    
    def build_code(node, prefix='', code={}):
        if node:
            if node.char is not None:
                code[node.char] = prefix
            build_code(node.left, prefix + '0', code)
            build_code(node.right, prefix + '1', code)
        return code
    
    root = heap[

0]
    return build_code(root)

# Example usage
text = "this is an example for huffman encoding"
freqs = Counter(text)
huffman_codes = huffman_coding(freqs)
print("Huffman Codes:", huffman_codes)
```

## **Explanation:**
- **Huffman Coding:** Builds a priority queue of nodes, merges the least frequent nodes, and generates the Huffman codes based on the tree structure.

