package main

import (
	"fmt"
	"github.com/KKrusti/mowers/ui"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Missing program arguments. Remember to attach the path to the file")
	} else {
		fileName := os.Args[1]
		console := ui.Console{}
		os.Exit(console.ConsoleExit(os.Stdout, fileName))
	}
}
