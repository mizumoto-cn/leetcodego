/**
34. Find First and Last Position of Element in Sorted Array
Medium

Given an array of integers nums sorted in ascending order,
find the starting and ending position of a given target value.
If target is not found in the array, return [-1, -1].
You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Example 3:
Input: nums = [], target = 0
Output: [-1,-1]
*/

package q34

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var _ = new(utils.Util)

// This may downgrade to O(n) in case of all elements are the same as it will cost O(n) to find the start and end.
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	var (
		low, high, mid int
		found          bool
	)

	low = 0
	high = len(nums) - 1
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == target {
			found = true
			break
		} else if nums[mid] < target {
			low = mid + 1
		} else { // nums[mid] > target
			high = mid - 1
		}
	}

	if !found {
		return []int{-1, -1}
	}

	var (
		start, end int
	)

	start = mid
	for start > 0 && nums[start-1] == target {
		start--
	}

	end = mid
	for end < len(nums)-1 && nums[end+1] == target {
		end++
	}

	return []int{start, end}
}

// Stable 2*O(log n) solution, will cost almost twice time in most cases.
func searchRange2(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1 // avoid overflow
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func findLast(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1 // avoid overflow
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

func Solve(nums []int, target int) []int {
	return searchRange(nums, target)
}

func Solve2(nums []int, target int) []int {
	return searchRange2(nums, target)
}
