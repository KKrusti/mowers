package infrastructure

import (
	"github.com/KKrusti/mowers/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewMowerCommand(t *testing.T) {
	command := []string{"1 2 N", "LMLMLMLMM"}
	got := NewMowerCommand(command)

	expected := MowerCommand{
		mowerCommandInitialConfig: valueobjects.MowerInitialConfig{
			Direction: "N",
			PosX:      1,
			PosY:      2,
		},
		mowerCommandMovements: []string{"L", "M", "L", "M", "L", "M", "L", "M", "M"},
	}

	assert.Equal(t, expected, got)
}

func TestMowerCommand_GetInitialStatus(t *testing.T) {
	command := []string{"1 2 N", "LMLMLMLMM"}
	mowerCommand := NewMowerCommand(command)

	got := mowerCommand.GetMowerInitialStatusCommand()

	expectedInitialStatus := valueobjects.MowerInitialConfig{
		Direction: "N",
		PosX:      1,
		PosY:      2,
	}
	assert.Equal(t, expectedInitialStatus, got)
}

func TestMowerCommand_GetMowerMoves(t *testing.T) {
	command := []string{"1 2 N", "LMLMLMLMM"}
	mowerCommand := NewMowerCommand(command)

	got := mowerCommand.GetMowerMoves()

	expectedMoves := []string{"L", "M", "L", "M", "L", "M", "L", "M", "M"}
	assert.Equal(t, expectedMoves, got)
}
