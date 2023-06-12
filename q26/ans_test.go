package q26_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q26"

	assert "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		nums   []int
		expect []int
		k      int
	}{
		{[]int{1, 1, 2}, []int{1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
	}

	for _, c := range cases {
		assert.Equal(t, c.k, question.Solve(c.nums))
		assert.Equal(t, c.expect, c.nums[:c.k])
	}
}
