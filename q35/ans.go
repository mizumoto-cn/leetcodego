/**
35. Search Insert Position
Easy

Given a sorted array of distinct integers and a target value,
return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4

Constraints:

1 <= nums.length <= 10^4
-10^4 <= nums[i] <= 10^4
nums contains distinct values sorted in ascending order.
-10^4 <= target <= 10^4

*/

package q35

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var _ = new(utils.Util)

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		// switch {
		// case nums[mid] == target || (mid == 0 && nums[mid] > target) || (mid > 0 && nums[mid-1] < target && nums[mid] > target):
		// 	return mid
		// case mid == len(nums)-1 && nums[mid] < target:
		// 	return mid + 1
		// case nums[mid] > target:
		// 	r = mid - 1
		// default:
		// 	l = mid + 1
		// }
		if nums[mid] >= target {
			r = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] >= target {
				return mid + 1
			}
			l = mid + 1
		}
	}
	return l
}

func Solve(nums []int, target int) int {
	return searchInsert(nums, target)
}
