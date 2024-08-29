# **Bucket Sort**

## **Overview**

Bucket Sort is a comparison-based algorithm that divides the array into a number of buckets, sorts each bucket individually (often using another sorting algorithm), and then concatenates the sorted buckets.

## **Where is it Used?**

- **Uniformly distributed data:** It’s effective when input data is uniformly distributed.
- **Floating-point numbers:** Often used for sorting floating-point numbers.

## **How Does it Work?**

Bucket Sort divides the array into a number of buckets, distributes the elements into these buckets, sorts each bucket, and then concatenates the sorted buckets.

## **Python Example**

```python
def bucket_sort(arr):
    bucket_count = len(arr)
    buckets = [[] for _ in range(bucket_count)]

    for num in arr:
        index = int(bucket_count * num)
        buckets[index].append(num)

    for i in range(bucket_count):
        buckets[i].sort()

    return [num for bucket in buckets for num in bucket]

# Test the function
arr = [0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68]
print(bucket_sort(arr))  # Output: [0.12, 0.17, 0.21, 0.23, 0.26, 0.39, 0.68, 0.72, 0.78, 0.94]
```

- **Explanation:**  
  The `bucket_sort` function distributes the elements into buckets, sorts each bucket, and then concatenates the sorted buckets into the final sorted array.
