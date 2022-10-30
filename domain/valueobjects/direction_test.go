package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_new_direction(t *testing.T) {
	orientation := "N"

	direction := NewDirection(orientation)

	assert.Equal(t, orientation, direction.CardinalPoint.String())
}
