package q1_test

import (
	"testing"

	"github.com/mizumoto-cn/leetcodego/q1"

	assert "github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	var cases = []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, q1.Solve(c.nums, c.target))
	}
}
