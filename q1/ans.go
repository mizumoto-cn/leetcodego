/**
a template q1
*/

package q1

import (
	"github.com/mizumoto-cn/leetcodego/util"
)

var _ = new(util.Util)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		expected := target - v
		if a, ok := m[expected]; ok {
			return []int{a, k}
		}
		m[v] = k
	}
	return nil
}

func Solve(nums []int, target int) []int {
	return twoSum(nums, target)
}
