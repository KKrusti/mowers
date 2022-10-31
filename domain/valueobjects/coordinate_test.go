package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_new_coordinate(t *testing.T) {
	xCoord := 3
	yCoord := 4

	newCoord := NewCoordinates(xCoord, yCoord)

	assert.Equal(t, xCoord, newCoord.X)
	assert.Equal(t, yCoord, newCoord.Y)
}

func Test_move_up(t *testing.T) {
	coordinates := NewCoordinates(1, 1)

	coordinates.MoveUp()
	expectedCoordinates := NewCoordinates(1, 2)

	assert.Equal(t, expectedCoordinates, coordinates)
}

func Test_move_down(t *testing.T) {
	coordinates := NewCoordinates(1, 1)

	coordinates.MoveDown()
	expectedCoordinates := NewCoordinates(1, 0)

	assert.Equal(t, expectedCoordinates, coordinates)
}

func Test_move_left(t *testing.T) {
	coordinates := NewCoordinates(1, 1)

	coordinates.MoveLeft()
	expectedCoordinates := NewCoordinates(0, 1)

	assert.Equal(t, expectedCoordinates, coordinates)
}

func Test_move_right(t *testing.T) {
	coordinates := NewCoordinates(1, 1)

	coordinates.MoveRight()
	expectedCoordinates := NewCoordinates(2, 1)

	assert.Equal(t, expectedCoordinates, coordinates)
}

func TestCoordinate_UP(t *testing.T) {
	coordinates := NewCoordinates(1, 1)

	newCoordinate := coordinates.Up()
	expectedCoordinates := NewCoordinates(1, 2)

	assert.Equal(t, expectedCoordinates, &newCoordinate)
}

func TestCoordinate_Down(t *testing.T) {
	coordinates := NewCoordinates(1, 1)

	newCoordinate := coordinates.Down()
	expectedCoordinates := NewCoordinates(1, 0)

	assert.Equal(t, expectedCoordinates, &newCoordinate)
}

func TestCoordinate_Left(t *testing.T) {
	coordinates := NewCoordinates(1, 1)

	newCoordinate := coordinates.Left()
	expectedCoordinates := NewCoordinates(0, 1)

	assert.Equal(t, expectedCoordinates, &newCoordinate)
}

func TestCoordinate_Right(t *testing.T) {
	coordinates := NewCoordinates(1, 1)

	newCoordinate := coordinates.Right()
	expectedCoordinates := NewCoordinates(2, 1)

	assert.Equal(t, expectedCoordinates, &newCoordinate)
}
