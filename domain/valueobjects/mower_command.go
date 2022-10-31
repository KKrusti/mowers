package valueobjects

type MowerCommand struct {
	plateauCommand            string
	mowerCommandInitialStatus string
	mowerCommandMovements     string
}

func NewMowerCommand(command []string) MowerCommand {
	return MowerCommand{
		plateauCommand:            command[0],
		mowerCommandInitialStatus: command[1],
		mowerCommandMovements:     command[2],
	}
}

func (mowerCommand MowerCommand) GetPlateauCommand() string {
	return mowerCommand.plateauCommand
}

func (mowerCommand MowerCommand) GetMowerInitialStatusCommand() string {
	return mowerCommand.mowerCommandInitialStatus
}

func (mowerCommand MowerCommand) GetMowerMovementsCommand() string {
	return mowerCommand.mowerCommandMovements
}
