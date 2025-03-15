package main

import (
	"fmt"
)

func binarySearch(arr []int, targetVal int) int{
	left := 0
	right := len(arr) - 1

	for left <= right {
		midVal := left + (right - left) / 2

		if arr[midVal] = targetVal {
			return midVal
		} else if arr[midVal] < left {
			left = midVal + 1
		} else {
			right = midVal - 1
		}
	}

	return -1
}