/**
541
*/

package q541

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

// Runtime
// 0ms
// Beats 100.00%of users with Go

// Memory
// 3.40MB
// Beats 100.00%of users with Go

func reverseStr(s string, k int) string {
	ss := []byte(s)
	l := len(ss)
	for i := 0; i < l; i += 2 * k {
		if i+k > l {
			reverse(ss[i:])
		} else {
			reverse(ss[i : i+k])
		}
	}
	return string(ss)
}

func reverse(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

func Solve(s string, k int) string {
	return reverseStr(s, k)
}
