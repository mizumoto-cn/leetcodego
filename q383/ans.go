/**
383
*/

package q383

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, r := range magazine {
		m[r]++
	}
	for _, r := range ransomNote {
		if m[r] <= 0 {
			return false
		}
		m[r]--
	}
	return true
}

func Solve(r, m string) bool {
	return canConstruct(r, m)
}
