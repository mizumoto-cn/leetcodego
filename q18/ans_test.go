package q18_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q18"
	"github.com/mizumoto-cn/leetcodego/utils"

	assert "github.com/stretchr/testify/assert"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func TestSolve(t *testing.T) {
	var cases = []struct {
		nums   []int
		target int
		expect [][]int
	}{
		// {[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		// {[]int{0, 0, 0, 0}, 0, [][]int{{0, 0, 0, 0}}},
		{[]int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5}, 0, [][]int{{-1, 0, 0, 1}}},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.nums, c.target))
	}
}
