package main

func binarySearch(arr []int, target int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	low := 0
	high := len(arr) - 1
	for low <= high {
		middle := (low + high) >> 1
		if arr[middle] == target {
			return middle
		}
		if arr[middle] < target {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -1
}