package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewMowerCommand(t *testing.T) {
	command := []string{"1 2 N", "LMLMLMLMM"}
	got := NewMowerCommand(command)

	expected := MowerCommand{
		//mowerCommandInitialConfig: "1 2 N",
		mowerCommandMovements: []string{"L", "M", "L", "M", "L", "M", "L", "M", "M"},
	}

	assert.Equal(t, expected, got)
}

func Test_parseMovements(t *testing.T) {
	//command := []string{"1 2 N", "LMLMLMLMM"}
	//mowerCommand := parseMovements(command)

}
