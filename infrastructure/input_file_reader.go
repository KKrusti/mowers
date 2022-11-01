package infrastructure

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type InputFileReader struct {
	reader *bufio.Scanner
}

func NewInputFileReader(fileName string) (InputFileReader, error) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fmt.Errorf("loading input file %v", err))
		return InputFileReader{}, err
	}
	return InputFileReader{
		reader: bufio.NewScanner(inputFile),
	}, nil
}

func (inputFileReader InputFileReader) ReadFromFile() (InputCommand, error) {
	plateau, err := inputFileReader.readPlateauParameters()
	if err != nil {
		return InputCommand{}, err
	}

	var command []string
	command = append(command, plateau)
	err = inputFileReader.readMowers(&command)
	if err != nil {
		return InputCommand{}, err
	}

	return NewInputCommand(command), nil
}

func (inputFileReader InputFileReader) readPlateauParameters() (string, error) {
	inputFileReader.reader.Scan()
	if inputFileReader.reader.Err() != nil {
		return "", inputFileReader.reader.Err()
	}
	return inputFileReader.reader.Text(), nil
}

func (inputFileReader InputFileReader) readMowers(command *[]string) error {
	lines := 0
	for inputFileReader.reader.Scan() {
		lines++
		read := inputFileReader.reader.Text()
		*command = append(*command, read)
	}
	if inputFileReader.reader.Err() != nil || lines%2 != 0 {
		return errors.New("file content incorrect")
	}
	return nil
}
