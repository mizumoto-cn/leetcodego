package q904_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q904"

	assert "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		nums   []int
		expect int
	}{
		{[]int{1, 2, 1}, 3},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{1, 2, 3, 2, 2}, 4},
	}
	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.nums))
	}
}
