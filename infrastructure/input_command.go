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
	mowerCommand := NewMowerCommand([]string{command[1], command[2]})

	return InputCommand{
		plateauCommand: command[0],
		mowerCommand:   []MowerCommand{mowerCommand},
	}
}

func (inputCommand InputCommand) GetMowerCommand() []MowerCommand {
	return inputCommand.mowerCommand
}

func (inputCommand InputCommand) GetPlateauDimensions() (int, int) {
	commandStrings := strings.Split(inputCommand.plateauCommand, " ")

	return shared.StringToInt(commandStrings[0]), shared.StringToInt(commandStrings[1])
}
