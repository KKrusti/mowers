package domain

import (
	"fmt"
	"github.com/KKrusti/mowers/domain/valueobjects"
)

const MOVE = "M"

type Mower struct {
	direction   *valueobjects.Direction
	coordinates *valueobjects.Coordinates
}

func NewMower(dir string, x, y int) Mower {
	coordinates := valueobjects.Coordinates{
		X: x,
		Y: y,
	}

	direction := valueobjects.NewDirection(dir)

	return Mower{
		direction:   direction,
		coordinates: &coordinates,
	}
}

func (mower *Mower) ExecuteMovement(command string) {
	switch command {
	case "R":
		mower.direction.RotateRight()
	case "L":
		mower.direction.RotateLeft()
	case MOVE:
		mower.move()
	default:
		//in case the command is not valid the mower stays in the same place, facing same direction
	}
}

func (mower *Mower) CheckNextMoveCoordinate() valueobjects.Coordinates {
	switch *mower.direction {
	case valueobjects.N:
		return mower.coordinates.Up()
	case valueobjects.S:
		return mower.coordinates.Down()
	case valueobjects.E:
		return mower.coordinates.Right()
	case valueobjects.W:
		return mower.coordinates.Left()
	default:
		return *mower.coordinates
	}

}

func (mower *Mower) move() {
	switch *mower.direction {
	case valueobjects.N:
		mower.coordinates.MoveUp()
		return
	case valueobjects.S:
		mower.coordinates.MoveDown()
		return
	case valueobjects.E:
		mower.coordinates.MoveRight()
		return
	case valueobjects.W:
		mower.coordinates.MoveLeft()
		return
	default:
		//this will never happen since when wrong direction is added, it's defaulted to N
	}
}

func (mower *Mower) directionAsString() string {
	return mower.direction.String()
}

func (mower *Mower) CurrentStatus() string {
	direction := mower.direction.String()
	posX := mower.coordinates.X
	posY := mower.coordinates.Y

	return fmt.Sprintf("%d %d %s", posX, posY, direction)
}

func (mower *Mower) GetCoordinates() *valueobjects.Coordinates {
	return mower.coordinates
}
