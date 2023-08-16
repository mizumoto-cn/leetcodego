package q707_test

import (
	"fmt"
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q707"
	"github.com/mizumoto-cn/leetcodego/utils"

	assert "github.com/stretchr/testify/assert"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func TestSolve(t *testing.T) {
	obj := question.Constructor()
	obj.AddAtHead(1)
	s := question.MyLinkedListToSlice(&obj)
	obj.AddAtTail(3)
	s = question.MyLinkedListToSlice(&obj)
	fmt.Printf("s: %v\n", s)
	assert.Equal(t, 1, obj.Get(0))
	assert.Equal(t, 3, obj.Get(1))
	s = question.MyLinkedListToSlice(&obj)
	obj.AddAtIndex(1, 2)
	assert.Equal(t, 2, obj.Get(1))
	obj.DeleteAtIndex(1)
	assert.Equal(t, 3, obj.Get(1))
}
