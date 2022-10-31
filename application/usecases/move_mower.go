package usecases

import (
	"github.com/KKrusti/mowers/domain"
	"github.com/KKrusti/mowers/domain/valueobjects"
)

type MoveMower struct {
	Mower   domain.Mower
	Plateau domain.Plateau
}

func (moveMowerHandler MoveMower) ExecuteCommand(mowerCommand valueobjects.MowerCommand) {

}
