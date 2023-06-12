package q977_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q977"

	assert "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		nums   []int
		expect []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.nums))
	}
}
