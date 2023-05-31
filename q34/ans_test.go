package q34_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q34"

	assert "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.nums, c.target))
	}
}

func TestSolve2(t *testing.T) {
	var cases = []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve2(c.nums, c.target))
	}
}
