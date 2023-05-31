package q367_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q367"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		num    int
		expect bool
	}{
		{16, true},
		{14, false},
		{1, true},
		{0, true},
		{2, false},
	}
	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.num))
	}
}
