package superdigitt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperDigit(t *testing.T) {
	t.Run("TEST1", func(t *testing.T) {
		result := SuperDigit("9875", 4)
		assert.Equal(t, result, int32(8))
	})
	t.Run("TEST2", func(t *testing.T) {
		result := SuperDigit("9", 100000)
		assert.Equal(t, int32(9), result)
	})

}
