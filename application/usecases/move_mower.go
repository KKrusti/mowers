package usecases

import (
	"github.com/KKrusti/mowers/domain"
	"github.com/KKrusti/mowers/infrastructure"
)

type MoveMower struct {
	Mower   domain.Mower
	Plateau domain.Plateau
}

func (moveMowerHandler MoveMower) ExecuteCommand(mowerCommand infrastructure.InputCommand) []string {
	moveMowerHandler.Plateau = domain.NewPlateau(mowerCommand.GetPlateauDimensions())

	result := moveMowerHandler.processMowers(mowerCommand.GetMowerCommand())

	return result
}

func (moveMowerHandler MoveMower) processMowers(mowerCommands []infrastructure.MowerCommand) []string {
	result := make([]string, len(mowerCommands))
	for i := 0; i < len(mowerCommands); i++ {
		mowerInitialConfig := mowerCommands[i].GetMowerInitialStatusCommand()
		moveMowerHandler.Mower = domain.NewMower(mowerInitialConfig.Direction, mowerInitialConfig.PosX, mowerInitialConfig.PosY)
		for _, move := range mowerCommands[i].GetMowerMoves() {
			if isCommandMove(move) {
				nextCoordinate := moveMowerHandler.Mower.CheckNextMoveCoordinate()
				if moveMowerHandler.Plateau.IsAvailableCoordinate(nextCoordinate) {
					moveMowerHandler.Plateau.UpdateOccupiedCoordinate(*moveMowerHandler.Mower.Coordinates, nextCoordinate)
					moveMowerHandler.Mower.ExecuteMovement(move)
				}
			} else {
				moveMowerHandler.Mower.ExecuteMovement(move)
			}
		}
		result[i] = moveMowerHandler.Mower.CurrentStatus()
	}
	return result
}

func isCommandMove(move string) bool {
	return move == domain.MOVE
}
