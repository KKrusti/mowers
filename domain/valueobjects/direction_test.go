package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_new_direction(t *testing.T) {
	type args struct {
		direction *Direction
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		{
			name: "N",
			args: args{
				direction: NewDirection("N"),
			},
			want: N,
		},
		{
			name: "E",
			args: args{
				direction: NewDirection("E"),
			},
			want: E,
		},
		{
			name: "S",
			args: args{
				direction: NewDirection("S"),
			},
			want: S,
		},
		{
			name: "W",
			args: args{
				direction: NewDirection("W"),
			},
			want: W,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, &tt.want, tt.args.direction)
		})
	}
}

func Test_rotate_left(t *testing.T) {
	type args struct {
		direction Direction
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		{
			name: "N->W",
			args: args{
				direction: N,
			},
			want: W,
		},
		{
			name: "W->S",
			args: args{
				direction: W,
			},
			want: S,
		},
		{
			name: "S->E",
			args: args{
				direction: S,
			},
			want: E,
		},
		{
			name: "E->N",
			args: args{
				direction: E,
			},
			want: N,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.direction.RotateLeft()
			assert.Equal(t, tt.want, tt.args.direction)
		})
	}
}

func Test_rotate_right(t *testing.T) {
	type args struct {
		direction Direction
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		{
			name: "N->E",
			args: args{
				direction: N,
			},
			want: E,
		},
		{
			name: "E->S",
			args: args{
				direction: E,
			},
			want: S,
		},
		{
			name: "S->W",
			args: args{
				direction: S,
			},
			want: W,
		},
		{
			name: "W->N",
			args: args{
				direction: W,
			},
			want: N,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.direction.RotateRight()
			assert.Equal(t, tt.want, tt.args.direction)
		})
	}
}
