package domain

import (
	"github.com/KKrusti/mowers/domain/valueobjects"
	"golang.org/x/exp/slices"
)

type Plateau struct {
	DimensionX          int
	DimensionY          int
	occupiedCoordinates []valueobjects.Coordinates
}

func NewPlateau(x, y int) Plateau {
	return Plateau{
		DimensionX:          x,
		DimensionY:          y,
		occupiedCoordinates: []valueobjects.Coordinates{},
	}
}

func (plateau *Plateau) IsAvailableCoordinate(coordinate valueobjects.Coordinates) bool {
	return plateau.isInsideLimits(coordinate) && !plateau.isOccupiedCoordinate(coordinate)
}

func (plateau *Plateau) isOccupiedCoordinate(coordinate valueobjects.Coordinates) bool {
	return slices.Contains(plateau.occupiedCoordinates, coordinate)
}

func (plateau *Plateau) AddOccupiedCoordinate(coordinates valueobjects.Coordinates) {
	if plateau.isInsideLimits(coordinates) {
		plateau.occupiedCoordinates = append(plateau.occupiedCoordinates, coordinates)
	}
}

func (plateau *Plateau) UpdateOccupiedCoordinate(oldCoordinates, newCoordinates valueobjects.Coordinates) {
	for i, coordinate := range plateau.occupiedCoordinates {
		if coordinate.X == oldCoordinates.X && coordinate.Y == oldCoordinates.Y {
			plateau.occupiedCoordinates[i] = newCoordinates
		}
	}
}

func (plateau *Plateau) isInsideLimits(coordinates valueobjects.Coordinates) bool {
	return plateau.isCoordinateXInsideLimits(coordinates) &&
		plateau.isCoordinateYInsideLimits(coordinates)
}

func (plateau *Plateau) isCoordinateXInsideLimits(coordinate valueobjects.Coordinates) bool {
	return coordinate.X >= 0 && coordinate.X <= plateau.DimensionX
}

func (plateau *Plateau) isCoordinateYInsideLimits(coordinate valueobjects.Coordinates) bool {
	return coordinate.Y >= 0 && coordinate.Y <= plateau.DimensionY
}
