# **Trie (Prefix Tree)**

## **Overview**

A Trie, or Prefix Tree, is a data structure used for storing a dynamic set of strings, where the keys are usually strings. It is useful for tasks involving prefix lookups, such as autocomplete and spell checking.

## **Where is it Used?**

- **Autocomplete:** Suggesting possible completions for a prefix.
- **Spell Checking:** Checking the presence of words in a dictionary.
- **String Matching:** Efficiently finding prefixes and substrings.

## **How Does it Work?**

A Trie is constructed by inserting each string character by character into a tree-like structure. Each node represents a character, and paths through the tree represent words or prefixes.

## **Python Example**

Here’s an example of implementing a basic Trie:

```python
class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end_of_word = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
        node.is_end_of_word = True

    def search(self, word):
        node = self.root
        for char in word:
            if char not in node.children:
                return False
            node = node.children[char]
        return node.is_end_of_word

# Example usage
trie = Trie()
trie.insert("hello")
print("Is 'hello' in trie?", trie.search("hello"))  # Output: True
print("Is 'hell' in trie?", trie.search("hell"))    # Output: False
```

- **Explanation:**  
  The `Trie` class provides methods to insert and search for words. The `insert` method adds words to the Trie, while the `search` method checks if a word exists in the Trie.

