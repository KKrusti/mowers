package usecases

import (
	"github.com/KKrusti/mowers/infrastructure"
)

type MoveMowerInterface interface {
	ExecuteCommand(command infrastructure.InputCommand) string
}
