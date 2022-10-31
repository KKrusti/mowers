package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMowerCommand(t *testing.T) {
	plateauCommand := "5 5"
	mowerInitialStatusCommand := "1 2 N"
	mowerCommandMovements := "LMLMLMLMM"

	mowerCommand := NewInputCommand([]string{plateauCommand, mowerInitialStatusCommand, mowerCommandMovements})

	assert.Equal(t, plateauCommand, mowerCommand.plateauCommand)
	//assert.Equal(t, mowerInitialStatusCommand, mowerCommand.mowerCommandInitialConfig)
	//assert.Equal(t, mowerCommandMovements, mowerCommand.mowerCommandMovements)
}

func TestMowerCommand_GetPlateauCommand(t *testing.T) {
	//plateauCommand := "5 5"
	//
	//mowerCommand := NewInputCommand([]string{plateauCommand, "", ""})
	//assert.Equal(t, plateauCommand, mowerCommand.GetPlateauDimensions())
}

//func TestMowerCommand_GetMowerInitialStatusCommand(t *testing.T) {
//	mowerInitialStatusCommand := "1 2 N"
//
//	mowerCommand := NewInputCommand([]string{"", mowerInitialStatusCommand, ""})
//	assert.Equal(t, mowerInitialStatusCommand, mowerCommand.GetMowerInitialStatusCommand())
//}

//func TestMowerCommand_GetMowerMovementsCommand(t *testing.T) {
//	mowerCommandMovements := "LMLMLMLMM"
//
//	mowerCommand := NewInputCommand([]string{"", "", mowerCommandMovements})
//
//	assert.Equal(t, mowerCommandMovements, mowerCommand.GetMowerMovementsCommand())
//}

//func TestMowerCommand_SeveralMowersCommands(t *testing.T) {
//	plateauCommand := "5 5"
//	mowerInitialStatusCommand1 := "1 2 N"
//	mowerCommandMovements1 := "LMLMLMLMM"
//	mowerInitialStatusCommand2 := "3 4 E"
//	mowerCommandMovements2 := "MRMMRMMLM"
//
//	mowerCommand := NewInputCommand([]string{plateauCommand, mowerInitialStatusCommand1, mowerCommandMovements1, mowerInitialStatusCommand2, mowerCommandMovements2})
//
//	assert.Equal(t, plateauCommand, mowerCommand.GetMowerMovementsCommand())
//	assert.Equal(t, mowerInitialStatusCommand1, mowerCommand.GetMowerMovementsCommand())
//	assert.Equal(t, mowerCommandMovements1, mowerCommand.GetMowerMovementsCommand())
//	assert.Equal(t, mowerInitialStatusCommand2, mowerCommand.GetMowerMovementsCommand())
//	assert.Equal(t, mowerCommandMovements2, mowerCommand.GetMowerMovementsCommand())
//}
