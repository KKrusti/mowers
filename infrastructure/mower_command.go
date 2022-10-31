package infrastructure

import (
	"github.com/KKrusti/mowers/domain/valueobjects"
	"github.com/KKrusti/mowers/shared"
	"strings"
)

type MowerCommand struct {
	mowerCommandInitialConfig valueobjects.MowerInitialConfig
	mowerCommandMovements     []string
}

func NewMowerCommand(command []string) MowerCommand {
	initialStatus := parseInitialConfig(command[0])
	movements := parseMovements(command[1])
	return MowerCommand{
		mowerCommandInitialConfig: initialStatus,
		mowerCommandMovements:     movements,
	}
}

func parseInitialConfig(command string) valueobjects.MowerInitialConfig {
	initialConfig := strings.Split(command, " ")
	return valueobjects.MowerInitialConfig{
		Direction: initialConfig[0],
		PosX:      shared.StringToInt(initialConfig[1]),
		PosY:      shared.StringToInt(initialConfig[2]),
	}
}

func parseMovements(command string) []string {
	runes := []rune(command)
	movements := make([]string, len(runes))

	for i := 0; i < len(command); i++ {
		movements[i] = string(runes[i])
	}
	return movements
}

func (mowerCommand MowerCommand) GetMowerInitialStatusCommand() valueobjects.MowerInitialConfig {
	return mowerCommand.mowerCommandInitialConfig
}

func (mowerCommand MowerCommand) GetMowerMoves() []string {
	return mowerCommand.mowerCommandMovements
}
