# **Anagram Checking**

## **Overview**

Anagram Checking involves determining if two strings are anagrams, meaning they contain the same characters with the same frequencies.

## **Where is it Used?**

- **String Analysis:** Verifying anagram properties in text processing.
- **Data Validation:** Checking for anagram conditions in various applications.

## **How Does it Work?**

Anagram checking can be performed by comparing sorted versions of the strings or by counting character frequencies and comparing the results.

## **Python Example**

Here’s an example of checking if two strings are anagrams:

```python
def are_anagrams(str1, str2):
    return sorted(str1) == sorted(str2)

# Example usage
str1 = "listen"
str2 = "silent"
print("Are anagrams:", are_anagrams(str1, str2))  # Output: True
```

- **Explanation:**  
  The `are_anagrams` function sorts both strings and compares them. If they are equal, the strings are anagrams.

