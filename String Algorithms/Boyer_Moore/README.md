# **Boyer-Moore Majority Vote**

## **Overview:**

The Boyer-Moore Majority Vote algorithm finds the majority element in an array—an element that appears more than half of the times. It operates in linear time and constant space by maintaining a candidate element and a count.

## **Python Example:**

```python
def majority_vote(nums):
    count, candidate = 0, None
    for num in nums:
        if count == 0:
            candidate = num
        count += (1 if num == candidate else -1)
    return candidate

# Example usage
nums = [3, 2, 3]
print("Majority element:", majority_vote(nums))
```

## **Explanation:**
- **Boyer-Moore Majority Vote Algorithm:** Identifies the majority element in an array in linear time using a candidate and count mechanism.

