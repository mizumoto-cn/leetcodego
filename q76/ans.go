/**
76. Minimum Window Substring
Hard

Given two strings s and t of lengths m and n respectively, return the minimum window substring
of s such that every character in t (including duplicates) is included in the window.
If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

Example 1:
	Input: s = "ADOBECODEBANC", t = "ABC"
	Output: "BANC"
	Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Example 2:
	Input: s = "a", t = "a"
	Output: "a"
	Explanation: The entire string s is the minimum window.

Example 3:
	Input: s = "a", t = "aa"
	Output: ""
	Explanation: Both 'a's from t must be included in the window.
*/

package q76

import (
	"github.com/mizumoto-cn/leetcodego/utils"
)

var _ = new(utils.Util)

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}
	var ms, mt [64]int
	for i := 0; i < len(t); i++ {
		mt[t[i]-'A']++
	}
	result, tail, head, min := "", 0, 0, len(s)+1
	for tail < len(s) {
		if mt[s[tail]-'A'] > 0 {
			ms[s[tail]-'A']++
		}
		for head <= tail && isMatch(ms, mt) {
			if tail-head+1 < min {
				min = tail - head + 1
				result = s[head : tail+1]
			}
			if mt[s[head]-'A'] > 0 {
				ms[s[head]-'A']--
			}
			head++
		}
		tail++
	}
	return result
}

func isMatch(ms, mt [64]int) bool {
	for i := 0; i < 64; i++ {
		if ms[i] < mt[i] {
			return false
		}
	}
	return true
}

func Solve(s string, t string) string {
	return minWindow(s, t)
}
