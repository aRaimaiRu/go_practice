package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQuickSort(t *testing.T) {
	t.Run("TEST1", func(t *testing.T) {
		expect := []int{1, 2, 3, 4, 5, 6}
		actual := QuickSort([]int{6, 5, 4, 3, 2, 1})
		assert.Equal(t, expect, actual)
	})
	t.Run("TEST1", func(t *testing.T) {
		expect := []int{1, 2, 3, 4, 5, 6}
		actual := QuickSort([]int{3, 5, 2, 4, 1, 6})
		assert.Equal(t, expect, actual)
	})

}
