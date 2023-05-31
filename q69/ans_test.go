package q69_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q69"

	assert "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		x      int
		expect int
	}{
		{4, 2},
		{8, 2},
		{0, 0},
		{1, 1},
		{2, 1},
		{99, 9},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.x))
	}
}
