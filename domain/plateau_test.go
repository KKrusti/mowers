package domain

import (
	"github.com/KKrusti/mowers/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_new_plateau(t *testing.T) {
	xCoord := 5
	yCoord := 5

	newPlateau := NewPlateau(xCoord, yCoord)
	expectedPlateau := Plateau{
		DimensionX:          xCoord,
		DimensionY:          yCoord,
		occupiedCoordinates: []valueobjects.Coordinates{},
	}

	assert.Equal(t, expectedPlateau, newPlateau)
}

func TestPlateau_IsAvailableCoordinate(t *testing.T) {
	plateau := NewPlateau(5, 5)
	coordinate := *valueobjects.NewCoordinates(2, 2)
	isAvailable := plateau.IsAvailableCoordinate(coordinate)

	assert.Equal(t, true, isAvailable)
}

func TestPlateau_IsInsideLimits(t *testing.T) {
	xcoordinate := 2
	ycoordinate := 3
	plateau := NewPlateau(xcoordinate, ycoordinate)

	type args struct {
		coordinate valueobjects.Coordinates
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "lower bound",
			args: args{
				coordinate: *valueobjects.NewCoordinates(0, 0),
			},
			want: true,
		},
		{
			name: "upper bound",
			args: args{
				coordinate: *valueobjects.NewCoordinates(xcoordinate, ycoordinate),
			},
			want: true,
		},
		{
			name: "below x",
			args: args{
				coordinate: *valueobjects.NewCoordinates(-1, 0),
			},
			want: false,
		},
		{
			name: "below y",
			args: args{
				coordinate: *valueobjects.NewCoordinates(0, -1),
			},
			want: false,
		},
		{
			name: "above x",
			args: args{
				coordinate: *valueobjects.NewCoordinates(xcoordinate+1, 0),
			},
			want: false,
		},
		{
			name: "above y",
			args: args{
				coordinate: *valueobjects.NewCoordinates(0, ycoordinate+1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plateau.isInsideLimits(tt.args.coordinate)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPlateau_isOccupiedCoordinate(t *testing.T) {
	type fields struct {
		DimensionX         int
		DimensionY         int
		occupiedCoordinate []valueobjects.Coordinates
	}
	type args struct {
		coordinate valueobjects.Coordinates
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "no occupied coordinates",
			fields: fields{
				DimensionX:         5,
				DimensionY:         5,
				occupiedCoordinate: []valueobjects.Coordinates{},
			},
			args: args{
				*valueobjects.NewCoordinates(2, 2),
			},
			want: false,
		},
		{
			name: "occupied coordinates",
			fields: fields{
				DimensionX: 5,
				DimensionY: 5,
				occupiedCoordinate: []valueobjects.Coordinates{
					{
						X: 2,
						Y: 2,
					}},
			},
			args: args{
				*valueobjects.NewCoordinates(2, 2),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plateau := Plateau{
				DimensionX:          tt.fields.DimensionX,
				DimensionY:          tt.fields.DimensionY,
				occupiedCoordinates: tt.fields.occupiedCoordinate,
			}
			assert.Equal(t, tt.want, plateau.isOccupiedCoordinate(tt.args.coordinate))
		})
	}
}

func TestPlateau_AddOccupiedCoordinate(t *testing.T) {
	type fields struct {
		DimensionX         int
		DimensionY         int
		occupiedCoordinate []valueobjects.Coordinates
	}
	type args struct {
		coordinate valueobjects.Coordinates
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []valueobjects.Coordinates
	}{
		{
			name: "initial empty, add coordinates inside limits",
			fields: fields{
				DimensionX:         5,
				DimensionY:         5,
				occupiedCoordinate: []valueobjects.Coordinates{},
			},
			args: args{
				*valueobjects.NewCoordinates(2, 2),
			},
			want: []valueobjects.Coordinates{{2, 2}},
		},
		{
			name: "already has values, add coordinates inside limits",
			fields: fields{
				DimensionX: 5,
				DimensionY: 5,
				occupiedCoordinate: []valueobjects.Coordinates{
					{
						X: 2,
						Y: 2,
					}},
			},
			args: args{
				*valueobjects.NewCoordinates(3, 4),
			},
			want: []valueobjects.Coordinates{{2, 2}, {3, 4}},
		},
		{
			name: "already has values, add coordiantes outside limits",
			fields: fields{
				DimensionX: 5,
				DimensionY: 5,
				occupiedCoordinate: []valueobjects.Coordinates{
					{
						X: 2,
						Y: 2,
					}},
			},
			args: args{
				*valueobjects.NewCoordinates(3, 6),
			},
			want: []valueobjects.Coordinates{{2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plateau := Plateau{
				DimensionX:          tt.fields.DimensionX,
				DimensionY:          tt.fields.DimensionY,
				occupiedCoordinates: tt.fields.occupiedCoordinate,
			}
			plateau.AddOccupiedCoordinate(tt.args.coordinate)
			assert.Equal(t, tt.want, plateau.occupiedCoordinates)
		})
	}
}

func TestPlateau_UpdateOccupiedCoordinate(t *testing.T) {
	plateau := NewPlateau(5, 5)
	oldCoordinates := valueobjects.NewCoordinates(1, 1)
	plateau.AddOccupiedCoordinate(*oldCoordinates)
	newCoordinates := valueobjects.NewCoordinates(1, 2)

	plateau.UpdateOccupiedCoordinate(*oldCoordinates, *newCoordinates)

	assert.Equal(t, false, plateau.isOccupiedCoordinate(*oldCoordinates))
	assert.Equal(t, true, plateau.isOccupiedCoordinate(*newCoordinates))
}
