package infrastructure

import (
	"github.com/KKrusti/mowers/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMowerCommand(t *testing.T) {
	inputPlateau := "5 5"
	mowerInitialStatusCommand := "1 2 N"
	mowerCommandMovements := "LMLMLMLMM"

	inputCommand := NewInputCommand([]string{inputPlateau, mowerInitialStatusCommand, mowerCommandMovements})
	expectedMowerCommand := []MowerCommand{{
		mowerCommandInitialConfig: valueobjects.MowerInitialConfig{
			Direction: "N",
			PosX:      1,
			PosY:      2,
		},
		mowerCommandMovements: []string{"L", "M", "L", "M", "L", "M", "L", "M", "M"},
	},
	}

	assert.Equal(t, inputPlateau, inputCommand.plateauCommand)
	assert.Equal(t, expectedMowerCommand, inputCommand.mowerCommand)
}

func TestMowerCommand_GetPlateauDimensions(t *testing.T) {
	plateauCommand := "5 4"
	mowerInitialStatusCommand := "1 2 N"
	mowerCommandMovements := "LMLMLMLMM"

	inputCommand := NewInputCommand([]string{plateauCommand, mowerInitialStatusCommand, mowerCommandMovements})
	dimensionX, dimensionY := inputCommand.GetPlateauDimensions()

	assert.Equal(t, 5, dimensionX)
	assert.Equal(t, 4, dimensionY)
}

func TestMowerCommand_GetMowerCommand(t *testing.T) {
	plateauCommand := "5 4"

	inputCommand := NewInputCommand([]string{plateauCommand, "1 2 N", "LMLMLMLMM"})
	mowerCommand := inputCommand.GetMowerCommand()

	expectedMowerCommand := []MowerCommand{{
		mowerCommandInitialConfig: valueobjects.MowerInitialConfig{
			Direction: "N",
			PosX:      1,
			PosY:      2,
		},
		mowerCommandMovements: []string{"L", "M", "L", "M", "L", "M", "L", "M", "M"},
	},
	}

	assert.Equal(t, expectedMowerCommand, mowerCommand)
}
