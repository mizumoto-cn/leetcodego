package q344_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q344"
	"github.com/mizumoto-cn/leetcodego/utils"

	assert "github.com/stretchr/testify/assert"
)

var util utils.Util

func init(){
	util = utils.NewUtil()
}

func TestSolve(t *testing.T) {
	var cases = []struct {
		// TODO: fill in your test case definition
	}{
		{ /* TODO: fill in your test cases */ },
		{ /* TODO: fill in your test cases */ },
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.nums, c.target))
	}
}