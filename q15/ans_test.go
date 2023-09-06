package q15_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q15"
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
		expect [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.nums))
	}
}
