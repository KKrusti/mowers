package domain

import (
	"fmt"
	"github.com/KKrusti/mowers/domain/valueobjects"
)

const MOVE = "M"

type Mower struct {
	Direction   *valueobjects.Direction
	Coordinates *valueobjects.Coordinates
}

func NewMower(dir string, x, y int) Mower {
	coordinates := valueobjects.Coordinates{
		X: x,
		Y: y,
	}

	direction := valueobjects.NewDirection(dir)

	return Mower{
		Direction:   direction,
		Coordinates: &coordinates,
	}
}

func (mower *Mower) ExecuteMovement(command string) {
	switch command {
	case "R":
		mower.Direction.RotateRight()
	case "L":
		mower.Direction.RotateLeft()
	case MOVE:
		mower.move()
	default:
		//in case the command is not valid the mower stays in the same place, facing same direction
	}
}

func (mower *Mower) CheckNextMoveCoordinate() valueobjects.Coordinates {
	switch *mower.Direction {
	case valueobjects.N:
		return mower.Coordinates.Up()
	case valueobjects.S:
		return mower.Coordinates.Down()
	case valueobjects.E:
		return mower.Coordinates.Right()
	case valueobjects.W:
		return mower.Coordinates.Left()
	default:
		return *mower.Coordinates
	}

}

func (mower *Mower) move() {
	switch *mower.Direction {
	case valueobjects.N:
		mower.Coordinates.MoveUp()
		return
	case valueobjects.S:
		mower.Coordinates.MoveDown()
		return
	case valueobjects.E:
		mower.Coordinates.MoveRight()
		return
	case valueobjects.W:
		mower.Coordinates.MoveLeft()
		return
	default:
		//this will never happen since when wrong direction is added, it's defaulted to N
	}
}

func (mower *Mower) directionAsString() string {
	return mower.Direction.String()
}

func (mower *Mower) CurrentStatus() string {
	direction := mower.Direction.String()
	posX := mower.Coordinates.X
	posY := mower.Coordinates.Y

	return fmt.Sprintf("%s %d %d", direction, posX, posY)
}
