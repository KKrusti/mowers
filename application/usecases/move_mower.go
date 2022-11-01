package usecases

import (
	"github.com/KKrusti/mowers/domain"
	"github.com/KKrusti/mowers/domain/valueobjects"
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
		currentCoordinate := moveMowerHandler.getMowerCurrentCoordinate()
		moveMowerHandler.Plateau.AddOccupiedCoordinate(*currentCoordinate)
		for _, move := range mowerCommands[i].GetMowerMoves() {
			if isCommandMove(move) {
				nextCoordinate := moveMowerHandler.Mower.CheckNextMoveCoordinate()
				if moveMowerHandler.isPlateauAvailableCoordinate(nextCoordinate) {
					moveMowerHandler.updatePlateauOccupiedCoordinate(nextCoordinate)
					moveMowerHandler.moveMower(move)
				}
			} else {
				moveMowerHandler.moveMower(move)
			}
		}
		result[i] = moveMowerHandler.getMowerCurrentStatus()
	}
	return result
}

func (moveMowerHandler MoveMower) isPlateauAvailableCoordinate(coordinate valueobjects.Coordinates) bool {
	return moveMowerHandler.Plateau.IsAvailableCoordinate(coordinate)
}

func (moveMowerHandler MoveMower) updatePlateauOccupiedCoordinate(nextCoordinate valueobjects.Coordinates) {
	moveMowerHandler.Plateau.UpdateOccupiedCoordinate(*moveMowerHandler.Mower.GetCoordinates(), nextCoordinate)
}

func (moveMowerHandler MoveMower) moveMower(move string) {
	moveMowerHandler.Mower.ExecuteMovement(move)
}

func (moveMowerHandler MoveMower) getMowerCurrentStatus() string {
	return moveMowerHandler.Mower.CurrentStatus()
}

func (moveMowerHandler MoveMower) getMowerCurrentCoordinate() *valueobjects.Coordinates {
	return moveMowerHandler.Mower.GetCoordinates()
}

func isCommandMove(move string) bool {
	return move == domain.MOVE
}
