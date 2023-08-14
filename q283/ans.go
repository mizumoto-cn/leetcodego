/**
283. Move Zeroes
Easy

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Example 1:
	Input: nums = [0,1,0,3,12]
	Output: [1,3,12,0,0]

Example 2:
	Input: nums = [0]
	Output: [0]
*/

package q283

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var _ = new(utils.Util)

// func moveZeroes(nums []int) {
// 	k := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != 0 {
// 			nums[k] = nums[i]
// 			k++
// 		}
// 	}
// 	for i := k; i < len(nums); i++ {
// 		nums[i] = 0
// 	}
// }

// faster: always swap non-zero element with zero element
func moveZeroes(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}
}

func Solve(nums []int) []int {
	moveZeroes(nums)
	return nums
}
