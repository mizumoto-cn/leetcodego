package q704_test

import (
	"testing"

	"github.com/mizumoto-cn/leetcodego/q704"

	assert "github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	var cases = []struct {
		nums   []int
		target int
		expect int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, q704.Solve(c.nums, c.target))
	}
}
