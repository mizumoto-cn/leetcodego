package q283_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q283"

	assert "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		nums   []int
		expect []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.nums))
	}
}
