package ui

import (
	"fmt"
	"github.com/KKrusti/mowers/application/usecases"
	"github.com/KKrusti/mowers/infrastructure"
	"io"
)

type Console struct {
	useCase usecases.MoveMowerInterface
}

func (console Console) ConsoleExit(out io.Writer, filePath string) int {
	fileReader, err := createFileReader(filePath)
	if err != nil {
		return 1
	}

	inputCommand, err := fileReader.ReadFromFile()
	if err != nil {
		println(fmt.Sprintf("error ocurred while reading from file: %v", err))
		return 1
	}

	console.useCase = usecases.MoveMower{}
	response := console.useCase.ExecuteCommand(inputCommand)
	_, err = fmt.Fprintf(out, response)
	if err != nil {
		println(fmt.Sprintf("error ocurred while reading writting to Standard output %v", err))
	}

	return 0
}

func createFileReader(filePath string) (infrastructure.InputFileReader, error) {
	fileReader, err := infrastructure.NewInputFileReader(filePath)
	if err != nil {
		println(fmt.Sprintf("error ocurred while opening the file %v", err))
		return infrastructure.InputFileReader{}, err
	}
	return fileReader, nil
}
