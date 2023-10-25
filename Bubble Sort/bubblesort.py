def bubblesort(arr):
    n = len(arr)

    for i in range(n):
        for j in range(0, n-i-1): # ostatnie i elementy są na miejscu
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = True
        if (swapped == False):
            break