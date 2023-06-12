/**
977. Squares of a Sorted Array
Easy

Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:
	Input: nums = [-4,-1,0,3,10]
	Output: [0,1,9,16,100]
	Explanation: After squaring, the array becomes [16,1,0,9,100].
	After sorting, it becomes [0,1,9,16,100].

Example 2:
	Input: nums = [-7,-3,2,3,11]
	Output: [4,9,9,49,121]
	Explanation: After squaring, the array becomes [49,9,4,9,121].
	After sorting, it becomes [4,9,9,49,121].
*/

package q977

import (
	"github.com/mizumoto-cn/leetcodego/util"
)

var _ = new(util.Util)

func sortedSquares(nums []int) []int {
	var (
		h, t = 0, len(nums) - 1
	)
	result := make([]int, len(nums))
	for h <= t {
		i := t - h
		if sq(nums[h]) > sq(nums[t]) {
			result[i] = sq(nums[h])
			h++
		} else {
			result[i] = sq(nums[t])
			t--
		}
	}
	return result
}

func sortedSquares_2(A []int) []int {
	ans := make([]int, len(A))
	for i, k, j := 0, len(A)-1, len(ans)-1; i <= j; k-- {
		if A[i]*A[i] > A[j]*A[j] {
			ans[k] = A[i] * A[i]
			i++
		} else {
			ans[k] = A[j] * A[j]
			j--
		}
	}
	return ans
}

func sq(n int) int { return n * n }

func Solve(nums []int) []int {
	return sortedSquares(nums)
}
