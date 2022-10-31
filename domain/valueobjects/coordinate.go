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

func (coordinates Coordinates) Up() Coordinates {
	return *NewCoordinates(coordinates.X, coordinates.Y+1)
}

func (coordinates Coordinates) Down() Coordinates {
	return *NewCoordinates(coordinates.X, coordinates.Y-1)
}

func (coordinates Coordinates) Left() Coordinates {
	return *NewCoordinates(coordinates.X-1, coordinates.Y)
}

func (coordinates Coordinates) Right() Coordinates {
	return *NewCoordinates(coordinates.Y+1, coordinates.Y)
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
