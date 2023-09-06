package q24_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q24"
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
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.in))
	}
}
