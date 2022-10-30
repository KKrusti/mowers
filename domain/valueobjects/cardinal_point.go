package valueobjects

type CardinalPoint int

const (
	N CardinalPoint = iota
	E
	S
	W
)

func (c CardinalPoint) String() string {
	return [...]string{"N", "E", "S", "W"}[c]
}

func (c *CardinalPoint) RotateRight() {
	if *c == 3 {
		*c = N
		return
	}
	*c = *c + 1
}

func (c *CardinalPoint) RotateLeft() {
	if *c == 0 {
		*c = W
		return
	}
	*c = *c - 1
}
