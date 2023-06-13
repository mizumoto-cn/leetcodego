package q209_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q209"

	assert "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		nums   []int
		target int
		expect int
	}{
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
		{[]int{1, 4, 4}, 4, 1},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.nums, c.target))
	}
}
