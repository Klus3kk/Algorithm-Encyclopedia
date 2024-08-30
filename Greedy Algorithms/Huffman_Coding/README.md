# **Huffman Coding**

## **Overview:**

Huffman Coding is a lossless data compression algorithm that uses variable-length codes for encoding symbols based on their frequencies. More frequent symbols use shorter codes.

## **Python Example:**

```python
import heapq
from collections import Counter, defaultdict

class Node:
    def __init__(self, char, freq):
        self.char = char
        self.freq = freq
        self.left = None
        self.right = None

    def __lt__(self, other):
        return self.freq < other.freq

def huffman_coding(s):
    # Calculate frequency of each character
    frequency = Counter(s)
    
    # Build priority queue
    heap = [Node(char, freq) for char, freq in frequency.items()]
    heapq.heapify(heap)
    
    # Build Huffman Tree
    while len(heap) > 1:
        left = heapq.heappop(heap)
        right = heapq.heappop(heap)
        merged = Node(None, left.freq + right.freq)
        merged.left = left
        merged.right = right
        heapq.heappush(heap, merged)
    
    # Generate Huffman Codes
    def generate_codes(node, prefix='', codebook=defaultdict()):
        if node is not None:
            if node.char is not None:
                codebook[node.char] = prefix
            generate_codes(node.left, prefix + '0', codebook)
            generate_codes(node.right, prefix + '1', codebook)
        return codebook

    root = heap[0]
    return generate_codes(root)

# Example usage
text = "this is an example for huffman encoding"
codes = huffman_coding(text)
print("Huffman Codes:", codes)
```

## **Explanation:**
- **Frequency Calculation:** Characters are counted to determine their frequencies.
- **Priority Queue:** Nodes are built and merged based on frequency.
- **Code Generation:** Huffman codes are assigned based on the tree structure.

