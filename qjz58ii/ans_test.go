package qjz58ii_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/qjz58ii"
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
		n      int
		expect string
	}{
		{"abcdefg", 2, "cdefgab"},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.s, c.n))
	}
}
