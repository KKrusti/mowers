package domain

type Plateau struct {
	DimensionX int
	DimensionY int
}

func NewPlateau(x, y int) Plateau {
	return Plateau{
		DimensionX: x,
		DimensionY: y,
	}
}
