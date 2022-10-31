package domain

import (
	"github.com/KKrusti/mowers/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_new_mower(t *testing.T) {
	direction := "N"
	xCoord := 4
	yCoord := 3

	newMower := NewMower(direction, xCoord, yCoord)

	expectedMower := Mower{
		Direction: valueobjects.NewDirection(direction),
		Coordinates: &valueobjects.Coordinates{
			X: xCoord,
			Y: yCoord,
		},
	}

	assert.Equal(t, expectedMower, newMower)
}

func Test_rotations(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "right",
			args: args{
				command: "R",
			},
			want: "E",
		},
		{
			name: "left",
			args: args{
				command: "L",
			},
			want: "W",
		},
		{
			name: "move just checking rotation",
			args: args{
				command: "M",
			},
			want: "N",
		},
		{
			name: "invalid command",
			args: args{
				command: "X",
			},
			want: "N",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mower := NewMower("N", 0, 0)
			mower.ExecuteMovement(tt.args.command)
			assert.Equal(t, tt.want, mower.directionAsString())
		})
	}
}

func Test_move(t *testing.T) {
	type args struct {
		mower Mower
	}
	tests := []struct {
		name string
		args args
		want *valueobjects.Coordinates
	}{
		{
			name: "up",
			args: args{
				mower: NewMower("N", 1, 1),
			},
			want: valueobjects.NewCoordinates(1, 2),
		},
		{
			name: "down",
			args: args{
				mower: NewMower("S", 1, 1),
			},
			want: valueobjects.NewCoordinates(1, 0),
		},
		{
			name: "left",
			args: args{
				mower: NewMower("W", 1, 1),
			},
			want: valueobjects.NewCoordinates(0, 1),
		},
		{
			name: "right",
			args: args{
				mower: NewMower("E", 1, 1),
			},
			want: valueobjects.NewCoordinates(2, 1),
		},
		{
			name: "Wrong value defaults to N",
			args: args{
				mower: NewMower("X", 1, 1),
			},
			want: valueobjects.NewCoordinates(1, 2),
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.args.mower.move()
			assert.Equal(t, tt.want, tt.args.mower.Coordinates)
		})
	}

}
