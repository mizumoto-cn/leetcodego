package q349_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q349"
	"github.com/mizumoto-cn/leetcodego/utils"

	assert "github.com/stretchr/testify/assert"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func TestSolve(t *testing.T) {
	var cases = []struct {
		m      []int
		n      []int
		expect []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.m, c.n))
	}
}
