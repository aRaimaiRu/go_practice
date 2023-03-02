package besttimetobuystock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperDigit(t *testing.T) {
	t.Run("Best Time To Buy Stock TEST1", func(t *testing.T) {
		expect := 5
		actual := maxProfit([]int{7, 1, 5, 3, 6, 4})
		assert.Equal(t, expect, actual)
	})
	t.Run("Best Time To Buy Stock TEST2", func(t *testing.T) {
		expect := 4
		actual := maxProfit([]int{3, 2, 6, 5, 0, 3})
		assert.Equal(t, expect, actual)
	})
	t.Run("Best Time To Buy Stock TEST3", func(t *testing.T) {
		expect := 2
		actual := maxProfit([]int{2, 1, 2, 1, 0, 1, 2})
		assert.Equal(t, expect, actual)
	})

}
