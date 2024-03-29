package pkg

const (
	ANS_TEMPLATE = `/**
{{.Name}}
*/
	
package q{{.Name}}
	
import (
	"github.com/mizumoto-cn/leetcodego/utils"
)
	
var util utils.Util

func init(){
	util = utils.NewUtil()
}
	
func /* TODO: Write your solution here */ {
	// TODO: Write your solution here
}
	
func Solve(/* TODO: fill in in/output type */ {
	return // TODO: call your solution here
}`

	ANS_TEST_TEMPLATE = `package q{{.Name}}_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q{{.Name}}"
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
}`
)
