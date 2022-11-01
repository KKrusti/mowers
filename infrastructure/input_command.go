package infrastructure

import (
	"github.com/KKrusti/mowers/shared"
	"strings"
)

type InputCommand struct {
	plateauCommand string
	mowerCommand   []MowerCommand
}

func NewInputCommand(command []string) InputCommand {
	var mowerCommandSlice []MowerCommand
	for i := 1; i < len(command); i = i + 2 {
		mowerCommands := []string{command[i], command[i+1]}
		mowerCommandSlice = append(mowerCommandSlice, NewMowerCommand(mowerCommands))
	}

	return InputCommand{
		plateauCommand: command[0],
		mowerCommand:   mowerCommandSlice,
	}
}

func (inputCommand InputCommand) GetMowerCommand() []MowerCommand {
	return inputCommand.mowerCommand
}

func (inputCommand InputCommand) GetPlateauDimensions() (int, int) {
	commandStrings := strings.Split(inputCommand.plateauCommand, " ")

	return shared.StringToInt(commandStrings[0]), shared.StringToInt(commandStrings[1])
}
