package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotate_left(t *testing.T) {
	type args struct {
		cardinalPoint CardinalPoint
	}
	tests := []struct {
		name string
		args args
		want CardinalPoint
	}{
		{
			name: "N->W",
			args: args{
				cardinalPoint: N,
			},
			want: W,
		},
		{
			name: "W->S",
			args: args{
				cardinalPoint: W,
			},
			want: S,
		},
		{
			name: "S->E",
			args: args{
				cardinalPoint: S,
			},
			want: E,
		},
		{
			name: "E->N",
			args: args{
				cardinalPoint: E,
			},
			want: N,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.cardinalPoint.RotateLeft()
			assert.Equal(t, tt.want, tt.args.cardinalPoint)
		})
	}
}

func Test_rotate_right(t *testing.T) {
	type args struct {
		cardinalPoint CardinalPoint
	}
	tests := []struct {
		name string
		args args
		want CardinalPoint
	}{
		{
			name: "N->E",
			args: args{
				cardinalPoint: N,
			},
			want: E,
		},
		{
			name: "E->S",
			args: args{
				cardinalPoint: E,
			},
			want: S,
		},
		{
			name: "S->W",
			args: args{
				cardinalPoint: S,
			},
			want: W,
		},
		{
			name: "W->N",
			args: args{
				cardinalPoint: W,
			},
			want: N,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.cardinalPoint.RotateRight()
			assert.Equal(t, tt.want, tt.args.cardinalPoint)
		})
	}
}
