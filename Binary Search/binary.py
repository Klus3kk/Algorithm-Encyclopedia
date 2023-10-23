# RECURSIVE
def binary_search(array,low,high,x):
    if high >= low:
        mid = (high + low) // 2

        if array[mid] == x:
            return mid
        elif array[mid] > x:
            return binary_search(array,low,mid - 1, x)
        else:
            return binary_search(array,mid+1,high,x)
    else:
        return -1
# ITERATIVE
def binary_search(array,x):
    low = 0
    high = len(array) - 1
    mid = 0
    while high >= low:
        mid = (low + high) // 2
        if array[mid] == x:
            return mid
        elif array[mid] < x:
            low = mid + 1
        else:
            high = mid - 1
    return -1