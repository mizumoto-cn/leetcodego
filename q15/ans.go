/**
15
*/

package q15

import (
	"fmt"
	"sort"

	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

// 32ms
// Beats 99.41%of users with Go
// 7.54MB
// Beats 76.20%of users with Go
func threeSum(nums []int) [][]int {
	// first sort the given array
	// maintain 3 pointers a b and c
	// a walks from the beginning to the end
	// while b/c walks from the head/tail to a
	// finding out possible solutions
	sort.Ints(nums)
	fmt.Printf("sorted: %v\n", nums)
	ans := [][]int{}
	l := len(nums)
	for a := 0; a < l; a++ {
		if nums[a] > 0 {
			return ans
		}
		b := a + 1
		c := l - 1
		for c > b {
			sum := nums[a] + nums[b] + nums[c]
			if sum > 0 {
				c--
			} else if sum < 0 {
				b++
			} else {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
				for c > b && nums[c] == nums[c-1] {
					c--
				}
				for c > b && nums[b] == nums[b+1] {
					b++
				}
				c--
				b++
			}
		}
		for a < l-2 && nums[a] == nums[a+1] {
			a++
		}
	}
	return ans
}

func Solve(nums []int) [][]int {
	ans := threeSum(nums)
	fmt.Printf("ans: %v\n", ans)
	return ans
}
