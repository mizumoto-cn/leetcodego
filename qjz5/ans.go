/**
jz5
*/

package qjz5

import (
	"fmt"

	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func replaceSpace(s string) string {
	// double pointer
	count := len(s)
	ss := []rune(s)
	fast := count - 1
	for _, c := range s {
		if c == ' ' {
			count++
		}
	}
	ans := ""
	for fast > -1 {
		if ss[fast] == ' ' {
			ans = fmt.Sprintf("%s%s", "%20", ans)
			fast--
		} else {
			ans = fmt.Sprintf("%c%s", ss[fast], ans)
			fast--
		}
	}
	return ans
}

func Solve(s string) string {
	return replaceSpace(s)
}
