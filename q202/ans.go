/**
202
*/

package q202

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

// 0ms 2.07mb
func isHappy(n int) bool {
	set := make(map[int]struct{})
	for n != 1 {
		if _, ok := set[n]; ok {
			return false
		}
		set[n] = struct{}{}
		n = getSum(n)
	}
	return true
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func Solve(n int) bool {
	return isHappy(n)
}
