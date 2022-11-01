package usecases

import (
	"github.com/KKrusti/mowers/domain"
	"github.com/KKrusti/mowers/domain/valueobjects"
	"github.com/KKrusti/mowers/infrastructure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveMower_ProcessMowers(t *testing.T) {
	type args struct {
		moveMower MoveMower
		commands  []infrastructure.MowerCommand
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "moves1",
			args: args{
				moveMower: MoveMower{
					Plateau: domain.NewPlateau(5, 5),
				},
				commands: []infrastructure.MowerCommand{
					infrastructure.NewMowerCommand([]string{"0 0 N", "MMRMMR"}),
				},
			},
			want: "2 2 S\n",
		},
		{
			name: "moves goes out of bounds",
			args: args{
				moveMower: MoveMower{
					Plateau: domain.NewPlateau(5, 5),
				},
				commands: []infrastructure.MowerCommand{
					infrastructure.NewMowerCommand([]string{"5 5 N", "MRMRMML"}),
				},
			},
			want: "5 3 E\n",
		},
		{
			name: "moves two mowers no collisions",
			args: args{
				moveMower: MoveMower{
					Plateau: domain.NewPlateau(5, 5),
				},
				commands: []infrastructure.MowerCommand{
					infrastructure.NewMowerCommand([]string{"1 1 N", "MMRMLM"}),
					infrastructure.NewMowerCommand([]string{"2 2 E", "RMMLMMMLL"}),
				},
			},
			want: "2 4 N\n5 0 W\n",
		},
		{
			name: "moves three mowers third collide with first",
			args: args{
				moveMower: MoveMower{
					Plateau: domain.NewPlateau(5, 5),
				},
				commands: []infrastructure.MowerCommand{
					infrastructure.NewMowerCommand([]string{"1 1 N", "MMRMLM"}),
					infrastructure.NewMowerCommand([]string{"2 2 E", "RMMLMMMLL"}),
					infrastructure.NewMowerCommand([]string{"1 5 S", "MLMMLM"}),
				},
			},
			want: "2 4 N\n5 0 W\n1 5 N\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.moveMower.processMowers(tt.args.commands)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMoveMower_isPlateauAvailableCoordinate(t *testing.T) {
	moveMower := MoveMower{
		Mower:   domain.NewMower("N", 1, 1),
		Plateau: domain.NewPlateau(4, 4),
	}

	occupiedCoordinates := valueobjects.NewCoordinates(2, 2)
	moveMower.Plateau.AddOccupiedCoordinate(*occupiedCoordinates)

	got := moveMower.isPlateauAvailableCoordinate(*occupiedCoordinates)

	assert.Equal(t, false, got)
}

func TestMoveMower_updateAvailableCoordinate(t *testing.T) {
	moveMower := MoveMower{
		Mower:   domain.NewMower("N", 2, 2),
		Plateau: domain.NewPlateau(4, 4),
	}
	occupiedCoordinates := valueobjects.NewCoordinates(2, 2)
	moveMower.Plateau.AddOccupiedCoordinate(*occupiedCoordinates)

	newCoordinate := valueobjects.NewCoordinates(2, 3)
	moveMower.updatePlateauOccupiedCoordinate(*newCoordinate)

	assert.Equal(t, false, moveMower.isPlateauAvailableCoordinate(*newCoordinate))
}

func TestMoveMower_moveMower(t *testing.T) {
	moveMower := MoveMower{
		Mower:   domain.NewMower("N", 2, 2),
		Plateau: domain.NewPlateau(4, 4),
	}

	moveMower.moveMower("M")
	expectedCoordiantes := valueobjects.NewCoordinates(2, 3)

	assert.Equal(t, expectedCoordiantes, moveMower.getMowerCurrentCoordinate())
}

func TestMoveMower_getCurrentStatus(t *testing.T) {
	moveMower := MoveMower{
		Mower:   domain.NewMower("N", 2, 2),
		Plateau: domain.NewPlateau(4, 4),
	}

	expectedCurrentStatus := "2 2 N"

	assert.Equal(t, expectedCurrentStatus, moveMower.getMowerCurrentStatus())
}

func TestMoveMower_getCurrentCoordinate(t *testing.T) {
	moveMower := MoveMower{
		Mower:   domain.NewMower("N", 2, 2),
		Plateau: domain.NewPlateau(4, 4),
	}

	expectedCoordinates := valueobjects.NewCoordinates(2, 2)

	assert.Equal(t, expectedCoordinates, moveMower.getMowerCurrentCoordinate())
}

func TestMoveMower_ExecuteCommand(t *testing.T) {
	moveMower := MoveMower{
		Mower:   domain.NewMower("N", 2, 2),
		Plateau: domain.NewPlateau(4, 4),
	}
	inputCommand := infrastructure.NewInputCommand([]string{"5 5", "1 2 N", "M"})

	got := moveMower.ExecuteCommand(inputCommand)

	expectedResult := "1 3 N\n"
	assert.Equal(t, expectedResult, got)
}
