package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_new_coordinate(t *testing.T) {
	xCoord := 3
	yCoord := 4

	newCoord := NewCoordinates(xCoord, yCoord)

	assert.Equal(t, xCoord, newCoord.X)
	assert.Equal(t, yCoord, newCoord.Y)
}
