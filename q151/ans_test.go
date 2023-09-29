package q151_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q151"
	"github.com/mizumoto-cn/leetcodego/utils"

	assert "github.com/stretchr/testify/assert"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func TestSolve(t *testing.T) {
	var cases = []struct {
		s      string
		expect string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world!  ", "world! hello"},
		{"a good   example", "example good a"},
		{"  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"Alice does not even like bob", "bob like even not does Alice"},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, question.Solve(c.s))
	}
}
