# **String Problems (Balanced Parentheses, Infix to Postfix Conversion)**

## **Overview**

String problems involving stacks and queues include checking for balanced parentheses and converting infix expressions to postfix (Reverse Polish Notation).

## **Where is it Used?**

- **Syntax Checking:** Validating expressions in compilers.
- **Expression Evaluation:** Converting and evaluating mathematical expressions.

## **How Does it Work?**

- **Balanced Parentheses:** Use a stack to ensure that each opening parenthesis has a corresponding closing parenthesis.
- **Infix to Postfix Conversion:** Use a stack to convert expressions from infix notation to postfix notation.

## **Python Example**

Here’s an example of checking balanced parentheses:

```python
def is_balanced(expression):
    stack = []
    matching = {')': '(', '}': '{', ']': '['}
    for char in expression:
        if char in matching.values():
            stack.append(char)
        elif char in matching.keys():
            if stack == [] or matching[char] != stack.pop():
                return False
    return stack == []

# Example usage
expression = "{[()()]}"
print("Balanced Parentheses:", is_balanced(expression))  # Output: True
```

- **Explanation:**  
  The `is_balanced` function checks if the parentheses in the string are balanced using a stack.

