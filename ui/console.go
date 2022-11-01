package ui

import (
	"fmt"
	"github.com/KKrusti/mowers/application/usecases"
	"github.com/KKrusti/mowers/infrastructure"
	"io"
)

const fileName = "input.txt"

func ConsoleExit(out io.Writer) int {
	fileReader, err := infrastructure.NewInputFileReader(fileName)
	if err != nil {
		println(fmt.Sprintf("error ocurred while opening the file %v", err))
		return 0
	}
	inputCommand, err := fileReader.ReadFromFile()
	if err != nil {
		println(fmt.Sprintf("error ocurred while reading from file %v", err))
		return 0
	}

	useCase := usecases.MoveMower{}
	response := useCase.ExecuteCommand(inputCommand)
	_, err = fmt.Fprintf(out, response)
	if err != nil {
		println(fmt.Sprintf("error ocurred while reading writting to Standard output %v", err))
	}
	return 0
}
