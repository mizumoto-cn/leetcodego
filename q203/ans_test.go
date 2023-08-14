package q203_test

import (
	"testing"

	question "github.com/mizumoto-cn/leetcodego/q203"
	"github.com/mizumoto-cn/leetcodego/utils"

	assert "github.com/stretchr/testify/assert"
)

var util utils.Util

func init() {
	util = utils.NewUtil()
}

func TestSolve(t *testing.T) {
	var cases = []struct {
		expect []int
		slice  []int
		val    int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 6, 3, 4, 5}, 6},
		{[]int{}, []int{}, 1},
		{[]int{}, []int{7, 7, 7, 7}, 7},
	}

	for _, c := range cases {
		assert.Equal(
			t,
			c.expect,
			util.LinkedListToSlice(
				question.Solve(util.SliceToLinkedList(c.slice), c.val),
			),
		)
	}
}
