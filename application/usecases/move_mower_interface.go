package usecases

import "github.com/KKrusti/mowers/domain/valueobjects"

type MoveMowerInterface interface {
	ExecuteCommand(command valueobjects.MowerCommand)
}
