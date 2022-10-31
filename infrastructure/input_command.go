package infrastructure

type InputCommand struct {
	plateauCommand string
	mowerCommand   []MowerCommand
}

func NewInputCommand(command []string) InputCommand {

	//mowerCommandsTotal := (len(command) - 1) / 2
	//mowerCommands := make([]string, mowerCommandsTotal)
	//for i := 0; i < mowerCommandsTotal; i = i + 2 {
	//	mowerCommand := NewMowerCommand(command[i+1:])
	//mowerCommands = append(mowerCommands, mowerCommand)
	//}

	return InputCommand{
		plateauCommand: command[0],
		//	mowerCommand:   NewMowerCommand(command[1:]),
	}
}

func (inputCommand InputCommand) GetMowerCommand() []MowerCommand {
	return inputCommand.mowerCommand
}

func (inputCommand InputCommand) GetPlateauDimensions() (int, int) {
	//return inputCommand.plateauCommand
	return 0, 0
}

func (inputCommand InputCommand) GetMowerInitialStatusCommand() string {
	//return inputCommand.mowerCommandInitialConfig
	return ""
}

func (inputCommand InputCommand) GetMowerMovementsCommand() string {
	//return inputCommand.mowerCommandMovements
	return ""
}
