package usecases

import (
	"github.com/KKrusti/mowers/domain"
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
		want []string
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
			want: []string{"S 2 2"},
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
			want: []string{"E 5 3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.moveMower.processMowers(tt.args.commands)
			assert.Equal(t, tt.want, got)
		})
	}
}
