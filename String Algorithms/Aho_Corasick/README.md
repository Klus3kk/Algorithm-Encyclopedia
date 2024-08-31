# **Aho-Corasick**

## **Overview:**

The Aho-Corasick algorithm is designed for searching multiple patterns simultaneously within a text. It uses a trie and failure links to efficiently handle multiple pattern matches in linear time relative to the size of the text and patterns.

## **Python Example:**

```python
from collections import deque, defaultdict

class AhoCorasick:
    def __init__(self):
        self.trie = defaultdict(dict)
        self.fail = {}
        self.output = defaultdict(list)
    
    def add_pattern(self, pattern):
        node = 0
        for char in pattern:
            if char not in self.trie[node]:
                self.trie[node][char] = len(self.trie)
            node = self.trie[node][char]
        self.output[node].append(pattern)
    
    def build_aho_corasick(self):
        queue = deque()
        for char in self.trie[0]:
            self.fail[self.trie[0][char]] = 0
            queue.append(self.trie[0][char])
        while queue:
            r = queue.popleft()
            for char in self.trie[r]:
                s = self.trie[r][char]
                queue.append(s)
                state = self.fail[r]
                while state is not None and char not in self.trie[state]:
                    state = self.fail.get(state)
                if state is None:
                    self.fail[s] = 0
                else:
                    self.fail[s] = self.trie[state][char]
                    self.output[s].extend(self.output.get(self.fail[s], []))
    
    def search(self, text):
        node = 0
        results = []
        for i, char in enumerate(text):
            while node is not None and char not in self.trie[node]:
                node = self.fail.get(node)
            if node is None:
                node = 0
                continue
            node = self.trie[node][char]
            if self.output[node]:
                for pattern in self.output[node]:
                    results.append((i - len(pattern) + 1, pattern))
        return results

# Example usage
ac = AhoCorasick()
ac.add_pattern("he")
ac.add_pattern("she")
ac.add_pattern("his")
ac.add_pattern("hers")
ac.build_aho_corasick()
text = "ushers"
print("Matches:", ac.search(text))
```

## **Explanation:**
- **Aho-Corasick Algorithm:** Efficiently searches for multiple patterns in a text by using a trie with failure links to handle multiple pattern matches in linear time.

