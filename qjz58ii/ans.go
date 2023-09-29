/**
jz58ii
*/

package qjz58ii

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func LeftRotateString(str string, n int) string {
	var (
		l = len(str)
		b = []byte(str)
	)
	if l == 0 {
		return ""
	}
	n = n % l
	// 反转区间为前n的子串
	// 反转区间为n到末尾的子串
	// 反转整个字符串
	b1 := reverse(b[0:n])
	b2 := reverse(b[n:l])
	return string(reverse(append(b1, b2...)))
}

func reverse(b []byte) []byte {
	l, r := 0, len(b)-1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return b
}

func LeftRotateString2(str string, n int) string {
	if len(str) == 0 {
		return str
	}
	n = n % len(str)
	return str[n:] + str[:n]
}

func Solve(str string, n int) string {
	return LeftRotateString2(str, n)
}
