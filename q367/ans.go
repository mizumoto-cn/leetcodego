/**
367. Valid Perfect Square
Easy

Given a positive integer num, return true if num is a perfect square or false otherwise.

You must not use any built-in exponent function or operator.

Example 1:
Input: num = 16
Output: true

Example 2:
Input: num = 14
Output: false
*/

package q367

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var _ = new(utils.Util)

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	mid := num >> 1
	for mid*mid > num {
		mid = (mid + num/mid) >> 1
	}
	return mid*mid == num
}

// Binary search
func isPerfectSquare2(num int) bool {
	l, h := 1, num
	for l <= h {
		mid := l + (h-l)>>1
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return false
}

func Solve(num int) bool {
	return isPerfectSquare(num)
}

func Solve2(num int) bool {
	return isPerfectSquare2(num)
}
