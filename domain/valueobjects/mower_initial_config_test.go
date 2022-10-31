package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewMowerInitialConfig(t *testing.T) {
	direction := "N"
	posX := 5
	posY := 5

	mowerInitialConfig := NewMowerInitialConfig(direction, posX, posY)

	expectedMowerConfig := MowerInitialConfig{
		Direction: direction,
		PosX:      posX,
		PosY:      posY,
	}
	assert.Equal(t, expectedMowerConfig, mowerInitialConfig)
}
