package valueobjects

type Direction struct {
	CardinalPoint *CardinalPoint
}

func NewDirection(direction string) Direction {
	var dir CardinalPoint
	switch direction {
	case "N":
		dir = N
	case "E":
		dir = E
	case "S":
		dir = S
	case "W":
		dir = W
	}

	return Direction{
		CardinalPoint: &dir,
	}
}

func (direction Direction) RotateRight() {
	direction.CardinalPoint.RotateRight()
}

func (direction Direction) RotateLeft() {
	direction.CardinalPoint.RotateLeft()
}
