package q19_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q19"
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
		n      int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4}},
		{[]int{1}, 1, []int{}},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.in, c.n))
	}
}
