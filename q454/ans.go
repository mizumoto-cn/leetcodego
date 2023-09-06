/**
454
*/

package q454

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mab := make(map[int]int)
	cnt := 0
	for _, a := range nums1 {
		for _, b := range nums2 {
			mab[a+b]++
		}
	}
	for _, c := range nums3 {
		for _, d := range nums4 {
			if v, ok := mab[-c-d]; ok {
				cnt += v
			}
		}
	}
	return cnt
}

func Solve(nums1, nums2, nums3, nums4 []int) int {
	return fourSumCount(nums1, nums2, nums3, nums4)
}
