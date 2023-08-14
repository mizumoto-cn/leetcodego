/**
27. Remove Element
Easy
*/

package q27

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var _ = new(utils.Util)

func removeElement(nums []int, val int) int {
	k := 0
	for _, n := range nums {
		if n != val {
			nums[k] = n
			k++
		}
	}
	return k
}

func Solve(nums []int, val int) int {
	return removeElement(nums, val)
}
