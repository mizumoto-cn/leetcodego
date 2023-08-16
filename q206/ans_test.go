package q206_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q206"
	"github.com/mizumoto-cn/leetcodego/utils"

	assert "github.com/stretchr/testify/assert"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func TestSolve(t *testing.T) {
	var cases = []struct {
		in     []int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1}, []int{1}},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.in))
	}
}
