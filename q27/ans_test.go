package q27_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q27"

	assert "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		nums   []int
		val    int
		expect int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{[]int{1}, 1, 0},
	}
	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.nums, c.val))
	}
}
