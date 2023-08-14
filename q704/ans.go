/**
Q704 binary-search

Easy

Given an array of integers nums which is sorted in ascending order, and an integer target,
write a function to search target in nums.

If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1

Constraints:

1 <= nums.length <= 10^4
-10^4 < nums[i], target < 10^4
All the integers in nums are unique.
nums is sorted in ascending order.
*/

package q704

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var _ = new(utils.Util)

// func search(nums []int, target int) int {
// 	if len(nums) == 0 {
// 		return -1
// 	}
// 	return binarySearch(nums, target, 0, len(nums)-1)
// }

// func binarySearch(nums []int, target int, left int, right int) int {
// 	if left > right {
// 		return -1
// 	}
// 	mid := (left + right) / 2
// 	if nums[mid] == target {
// 		return mid
// 	}
// 	if nums[mid] > target {
// 		return binarySearch(nums, target, left, mid-1)
// 	}
// 	return binarySearch(nums, target, mid+1, right)
// }

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1 // avoid overflow
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1 // mid is bigger than target, so r = mid - 1
		} else {
			l = mid + 1 // mid is smaller than target, so l = mid + 1
		}
	}
	return -1
}

func Solve(nums []int, target int) int {
	return search(nums, target)
}
