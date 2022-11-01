package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewInputFileReader(t *testing.T) {
	fileName := "../resources/testdata/input.txt"
	inputFileReader, err := NewInputFileReader(fileName)
	require.NoError(t, err)

	assert.Equal(t, nil, inputFileReader.reader.Err())
}

func TestNewInputFileReader_error(t *testing.T) {
	fileName := "../resources/testdata/inputt.txt"
	_, err := NewInputFileReader(fileName)
	require.Error(t, err)
}

func TestInputFileReader_ReadFromFile(t *testing.T) {
	fileName := "../resources/testdata/input.txt"
	inputFileReader, err := NewInputFileReader(fileName)
	require.NoError(t, err)

	inputCommand, err := inputFileReader.ReadFromFile()
	require.NoError(t, err)

	expectedInputCommand := NewInputCommand([]string{"5 5", "1 2 N", "LMLMLMLMM", "3 3 E", "MMRMMRMRRM"})
	assert.Equal(t, expectedInputCommand, inputCommand)
}

func TestInputFileReader_readPlateauParameters(t *testing.T) {
	fileName := "../resources/testdata/input_with_no_mowers.txt"
	inputFileReader, err := NewInputFileReader(fileName)
	require.NoError(t, err)

	got, err := inputFileReader.readPlateauParameters()
	require.NoError(t, err)

	expectedCommand := "5 5"
	assert.Equal(t, expectedCommand, got)
}

func TestInputFileReader_readMowers(t *testing.T) {
	fileName := "../resources/testdata/input_with_no_plateau.txt"
	inputFileReader, err := NewInputFileReader(fileName)
	require.NoError(t, err)
	command := []string{"5 5"}

	err = inputFileReader.readMowers(&command)
	require.NoError(t, err)

	expectedCommand := []string{"5 5", "1 2 N", "LMLMLMLMM", "3 3 E", "MMRMMRMRRM"}
	assert.Equal(t, expectedCommand, command)
}
