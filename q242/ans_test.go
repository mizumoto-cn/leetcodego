package q242_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q242"
	"github.com/mizumoto-cn/leetcodego/utils"

	assert "github.com/stretchr/testify/assert"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func TestSolve(t *testing.T) {
	var cases = []struct {
		s      string
		t      string
		expect bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.s, c.t))
	}
}
