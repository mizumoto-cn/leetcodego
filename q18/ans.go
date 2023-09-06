/**
18
*/

package q18

import (
	"sort"

	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	ans := [][]int{}
	for a := 0; a < l-3; a++ {
		for b := a + 1; b < l-2; b++ {
			c := b + 1
			d := l - 1
			for d > c {
				sum := nums[a] + nums[b] + nums[c] + nums[d] - target
				if sum > 0 {
					d--
				} else if sum < 0 {
					c++
				} else {
					ans = append(ans, []int{nums[a], nums[b], nums[c], nums[d]})
					for c < d && nums[c] == nums[c+1] {
						c++
					}
					for c < d && nums[d] == nums[d-1] {
						d--
					}
					c++
					d--
				}
			}
			for b < l-2 && nums[b] == nums[b+1] {
				b++
			}
		}
		for a < l-3 && nums[a] == nums[a+1] {
			a++
		}
	}
	return ans
}

func Solve(nums []int, target int) [][]int {
	return fourSum(nums, target)
}
