package qjz5_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/qjz5"
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
		expect string
	}{
		{"We are happy.", "We%20are%20happy."},
		{" ", "%20"},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.s), "input:%v", c)
	}
}
