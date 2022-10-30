package valueobjects

type Coordinates struct {
	X int
	Y int
}

func NewCoordinates(x, y int) *Coordinates {
	return &Coordinates{
		X: x,
		Y: y,
	}
}

func (coordinates *Coordinates) MoveUp() {
	coordinates.Y += 1
}

func (coordinates *Coordinates) MoveDown() {
	coordinates.Y -= 1
}

func (coordinates *Coordinates) MoveRight() {
	coordinates.X += 1
}

func (coordinates *Coordinates) MoveLeft() {
	coordinates.X -= 1
}
