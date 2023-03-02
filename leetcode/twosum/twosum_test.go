package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperDigit(t *testing.T) {
	t.Run("TEST1", func(t *testing.T) {
		expect := []int{0, 1}
		actual := TwoSum([]int{2, 7, 11, 15}, 9)
		assert.Equal(t, expect, actual)
	})

}
