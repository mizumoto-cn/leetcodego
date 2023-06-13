/**
209. Minimum Size Subarray Sum
Medium

Given an array of positive integers nums and a positive integer target,
return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr]
of which the sum is greater than or equal to target.
If there is no such subarray, return 0 instead.

Example 1:
	Input: target = 7, nums = [2,3,1,2,4,3]
	Output: 2
	Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Example 2:
	Input: target = 4, nums = [1,4,4]
	Output: 1

Example 3:
	Input: target = 11, nums = [1,1,1,1,1,1,1,1]
	Output: 0
*/

package q209

import (
	"github.com/mizumoto-cn/leetcodego/util"
)

var _ = new(util.Util)

func minSubArrayLen(target int, nums []int) int {
	var head, tail, sum, min int = 0, 0, 0, len(nums) + 1
	for tail < len(nums) {
		sum += nums[tail]
		for sum >= target {
			if (tail - head + 1) < min {
				min = tail - head + 1
			}
			sum -= nums[head]
			head++
		}
		tail++
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}

func minSubArrayLen_using_min(target int, nums []int) int {
	head, sum, ml := 0, 0, len(nums)+1
	for tail, v := range nums {
		sum += v
		for sum >= target {
			ml = min(ml, tail-head+1)
			sum -= nums[head]
			head++
		}
	}
	if ml == len(nums)+1 {
		return 0
	}
	return ml
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Solve(nums []int, target int) int {
	if target == (target/2)*2 {
		return minSubArrayLen(target, nums)
	}
	return minSubArrayLen_using_min(target, nums)
}
