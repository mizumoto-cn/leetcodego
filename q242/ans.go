/**
242
*/

package q242

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func isAnagram(s string, t string) bool {
	var m map[rune]int = make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for _, r := range t {
		m[r]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func Solve(s, t string) bool {
	return isAnagram(s, t)
}
