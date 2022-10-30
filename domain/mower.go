package domain

import "github.com/KKrusti/mowers/domain/valueobjects"

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

func (mower *Mower) ExecuteCommand(command string) {
	switch command {
	case "R":
		mower.Direction.RotateRight()
	case "L":
		mower.Direction.RotateLeft()
	case "M":
		mower.move()
	default:
		//in case the command is not valid the mower stays in the same place, facing same direction
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
		//do not move
	}

}

func (mower *Mower) directionAsString() string {
	return mower.Direction.String()
}
