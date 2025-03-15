package main

import (
	"fmt"
	"math"
)

// Note: Need to learn Binary search

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		// Partition nums1
		partition1 := (low + high) / 2
		// Partition nums2 (ensures left partition size equals (m + n + 1) / 2)
		partition2 := (m+n+1)/2 - partition1

		// Handle edge cases for nums1
		var maxLeft1, minRight1 int
		if partition1 == 0 {
			maxLeft1 = math.MinInt64 // Left partition of nums1 is empty
		} else {
			maxLeft1 = nums1[partition1-1]
		}
		if partition1 == m {
			minRight1 = math.MaxInt64 // Right partition of nums1 is empty
		} else {
			minRight1 = nums1[partition1]
		}

		// Handle edge cases for nums2
		var maxLeft2, minRight2 int
		if partition2 == 0 {
			maxLeft2 = math.MinInt64 // Left partition of nums2 is empty
		} else {
			maxLeft2 = nums2[partition2-1]
		}
		if partition2 == n {
			minRight2 = math.MaxInt64 // Right partition of nums2 is empty
		} else {
			minRight2 = nums2[partition2]
		}

		// Check if the partition is correct
		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			// Calculate median
			if (m+n)%2 == 0 {
				return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2
			} else {
				return float64(max(maxLeft1, maxLeft2))
			}
		} else if maxLeft1 > minRight2 {
			// Move partition1 to the left
			high = partition1 - 1
		} else {
			// Move partition1 to the right
			low = partition1 + 1
		}
	}

	// Should never reach here
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println("Median:", findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println("Median:", findMedianSortedArrays(nums1, nums2))
}