fun bubblesort(arr: IntArray) {
    val n = arr.size
    bool swapped = false
    for (i in 0 until n - 1) {
        for (j in 0 until n - i - 1) {
            if (arr[j] > arr[j + 1]) {
                val temp = arr[j]
                arr[j] = arr[j + 1]
                arr[j + 1] = temp
                swapped = true
            }
        }
        if (swapped == false) {
            break
        }
    }
}