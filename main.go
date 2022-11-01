package main

import (
	"github.com/KKrusti/mowers/ui"
	"os"
)

func main() {
	os.Exit(ui.ConsoleExit(os.Stdout))
}
