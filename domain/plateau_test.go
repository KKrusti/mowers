package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_new_plateau(t *testing.T) {
	xCoord := 5
	yCoord := 5

	newPlateau := NewPlateau(xCoord, yCoord)
	expectedPlateau := Plateau{
		DimensionX: xCoord,
		DimensionY: yCoord,
	}

	assert.Equal(t, expectedPlateau, newPlateau)
}
