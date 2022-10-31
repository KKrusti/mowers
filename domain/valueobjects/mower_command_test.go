package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMowerCommand(t *testing.T) {
	plateauCommand := "5 5"
	mowerInitialStatusCommand := "1 2 N"
	mowerCommandMovements := "LMLMLMLMM"

	mowerCommand := NewMowerCommand([]string{plateauCommand, mowerInitialStatusCommand, mowerCommandMovements})

	assert.Equal(t, plateauCommand, mowerCommand.plateauCommand)
	assert.Equal(t, mowerInitialStatusCommand, mowerCommand.mowerCommandInitialStatus)
	assert.Equal(t, mowerCommandMovements, mowerCommand.mowerCommandMovements)
}

func TestMowerCommand_GetPlateauCommand(t *testing.T) {
	plateauCommand := "5 5"

	mowerCommand := NewMowerCommand([]string{plateauCommand, "", ""})
	assert.Equal(t, plateauCommand, mowerCommand.GetPlateauCommand())
}
func TestMowerCommand_GetMowerInitialStatusCommand(t *testing.T) {
	mowerInitialStatusCommand := "1 2 N"

	mowerCommand := NewMowerCommand([]string{"", mowerInitialStatusCommand, ""})
	assert.Equal(t, mowerInitialStatusCommand, mowerCommand.GetMowerInitialStatusCommand())
}
func TestMowerCommand_GetMowerMovementsCommand(t *testing.T) {
	mowerCommandMovements := "LMLMLMLMM"

	mowerCommand := NewMowerCommand([]string{"", "", mowerCommandMovements})

	assert.Equal(t, mowerCommandMovements, mowerCommand.GetMowerMovementsCommand())
}
