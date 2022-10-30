package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_new_direction(t *testing.T) {
	type args struct {
		direction Direction
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "N",
			args: args{
				direction: NewDirection("N"),
			},
			want: "N",
		},
		{
			name: "E",
			args: args{
				direction: NewDirection("E"),
			},
			want: "E",
		},
		{
			name: "S",
			args: args{
				direction: NewDirection("S"),
			},
			want: "S",
		},
		{
			name: "W",
			args: args{
				direction: NewDirection("W"),
			},
			want: "W",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.args.direction.CardinalPoint.String())
		})
	}
}

func Test_rotations_right(t *testing.T) {
	type args struct {
		direction Direction
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "N->E",
			args: args{
				direction: NewDirection("N"),
			},
			want: "E",
		},
		{
			name: "E-S",
			args: args{
				direction: NewDirection("E"),
			},
			want: "S",
		},
		{
			name: "S",
			args: args{
				direction: NewDirection("S"),
			},
			want: "W",
		},
		{
			name: "W",
			args: args{
				direction: NewDirection("W"),
			},
			want: "N",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.direction.RotateRight()
			assert.Equal(t, tt.want, tt.args.direction.CardinalPoint.String())
		})
	}
}

func Test_rotations_left(t *testing.T) {
	type args struct {
		direction Direction
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "N->W",
			args: args{
				direction: NewDirection("N"),
			},
			want: "W",
		},
		{
			name: "E-N",
			args: args{
				direction: NewDirection("E"),
			},
			want: "N",
		},
		{
			name: "S",
			args: args{
				direction: NewDirection("S"),
			},
			want: "E",
		},
		{
			name: "W",
			args: args{
				direction: NewDirection("W"),
			},
			want: "S",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.direction.RotateLeft()
			assert.Equal(t, tt.want, tt.args.direction.CardinalPoint.String())
		})
	}
}
