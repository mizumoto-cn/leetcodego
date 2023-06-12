/**
844. Backspace String Compare
Easy

Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.
*/

package q844

import (
	"github.com/mizumoto-cn/leetcodego/util"
)

var _ = new(util.Util)

func backspaceCompare(s string, t string) bool {
	var (
		is     int    = 0
		it     int    = 0
		_s, _t []rune = make([]rune, len(s)), make([]rune, len(t))
	)
	for _, c := range s {
		if c == '#' {
			if is > 0 {
				is--
			}
		} else {
			_s[is] = c
			is++
		}
	}
	for _, c := range t {
		if c == '#' {
			if it > 0 {
				it--
			}
		} else {
			_t[it] = c
			it++
		}
	}
	return string(_s[:is]) == string(_t[:it])
}

func Solve(s string, t string) bool {
	return backspaceCompare(s, t)
}
