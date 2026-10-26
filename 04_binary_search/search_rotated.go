// Pattern: Binary Search on a rotated sorted array.
//
// Problem (LeetCode 33, Search in Rotated Sorted Array): an ascending array
// of distinct values was rotated at an unknown pivot. Find target's index
// or return -1, in O(log n).
//
// Idea: at every step, one half of the current window is properly sorted.
// If target lies inside the sorted half, search there; else search the
// other half. Decide by comparing against the window's ends.
package binarysearch

func SearchRotated(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] { // left half is sorted
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // right half is sorted
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
