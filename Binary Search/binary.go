package main

import "fmt"

func binary(arr [] int, x int) int {
	var low = 0
	var high = len(arr) - 1
	var mid = 0
	for high >= low {
		mid = (low + high) / 2
		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}