package valueobjects

type MowerInitialConfig struct {
	Direction string
	PosX      int
	PosY      int
}

func NewMowerInitialConfig(direction string, posX, posY int) MowerInitialConfig {
	return MowerInitialConfig{
		Direction: direction,
		PosX:      posX,
		PosY:      posY,
	}
}
