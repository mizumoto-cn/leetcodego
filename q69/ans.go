/**
69. Sqrt(x)
Easy

Given a non-negative integer x, compute and return the square root of x
rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.const

Example 1:
Input: x = 4
Output: 2

Example 2:
Input: x = 8
Output: 2
*/

package q69

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var _ = new(utils.Util)

// Newton's method: https://en.wikipedia.org/wiki/Integer_square_root#Algorithm_using_Newton's_method
// One way of calculating √n and isqrt(n) is to use Heron's method, which is a special case of Newton's method,
// to find a solution of the equation x^2 − n = 0, giving the iterative formula:
// x[k+1] = (x[k] + n/x[k]) / 2 , k ≥ 0, x[0] > 0
// The sequence x[k] converges quadratically to √n, as k → ∞.
func mySqrt(x int) int {
	// if x == 0 {
	// 	return 0
	// }
	if x == 1 {
		return 1
	}
	mid := x >> 1
	for mid*mid > x {
		mid = (mid + x/mid) >> 1
	}
	return mid
}

// Of course, we can use binary search to find the answer, but it will just be slower.

func Solve(x int) int {
	return mySqrt(x)
}
