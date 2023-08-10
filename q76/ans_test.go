package q76_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q76"

	assert "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		s   string
		t   string
		ans string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"a", "b", ""},
		{"a", "aa", ""},
		{"a", "aaa", ""},
	}
	for _, cas := range cases {
		assert.Equal(t, cas.ans, question.Solve(cas.s, cas.t))
	}
}
