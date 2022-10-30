package valueobjects

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

func (c *Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[*c]
}

func (c *Direction) RotateRight() {
	if *c == 3 {
		*c = N
		return
	}
	*c = *c + 1
}

func (c *Direction) RotateLeft() {
	if *c == 0 {
		*c = W
		return
	}
	*c = *c - 1
}

// NewDirection method set the direction the mower is facing to. In case a wrong value is given, it'll be placed facing North
func NewDirection(direction string) *Direction {
	var dir Direction
	switch direction {
	case "N":
		dir = N
	case "E":
		dir = E
	case "S":
		dir = S
	case "W":
		dir = W
	default:
		dir = N
	}

	return &dir
}

func (direction *Direction) asString() string {
	return direction.String()
}
