package q35_test

import (
	"testing"

	"github.com/mizumoto-cn/leetcodego/q35"
	assert "github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	var cases = []struct {
		nums   []int
		target int
		expect int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, q35.Solve(c.nums, c.target))
	}
}
