/**
349
*/

package q349

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for _, n := range nums1 {
		m[n] = struct{}{}
	}
	var res []int
	for _, n := range nums2 {
		if _, ok := m[n]; ok {
			res = append(res, n)
			delete(m, n)
		}
	}
	return res
}

func Solve(nums1, nums2 []int) []int {
	return intersection(nums1, nums2)
}
