function binary(arr,x) {
    var low = 0;
    var high = length(arr) - 1;
    var mid = 0;
    while (high >= low) {
        mid = (low + high) / 2;
        if (arr[mid] == x) {
            return mid;
        } else if (arr[mid] > x) {
            high = mid - 1;
        } else {
            low = mid + 1;
        }
    }
    return -1;   
}