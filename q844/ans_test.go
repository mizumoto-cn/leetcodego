package q844_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q844"

	assert "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var cases = []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "ab#c", t: "ad#c", expected: true},
		{s: "ab##", t: "c#d#", expected: true},
		{s: "a##c", t: "b", expected: false},
	}

	for _, tc := range cases {
		actual := question.Solve(tc.s, tc.t)
		assert.Equal(t, tc.expected, actual, "%+v", tc)
	}
}
