package main

import (
	"os"

	"github.com/mitchellh/cli"
)

var version = "0.1.0"

func main() {
	c := &cli.CLI{
		Name:       "jest-diff-ls",
		Version:    version,
		Args:       os.Args[1:],
		HelpWriter: os.Stdout,
	}

	ui := &cli.ColoredUi{
		ErrorColor: cli.UiColorRed,
		WarnColor:  cli.UiColorYellow,
		Ui: &cli.BasicUi{
			Writer:      os.Stdout,
			Reader:      os.Stdin,
			ErrorWriter: os.Stderr,
		},
	}

	c.Commands = map[string]cli.CommandFactory{}

	exitStatus, err := c.Run()

	if err != nil {
		ui.Error("Error: " + err.Error())
	}

	os.Exit(exitStatus)
}
