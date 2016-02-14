package main

import (
	"github.com/mitchellh/cli"
	"github.com/webcanvas/pinch/commands"
	"github.com/webcanvas/pinch/ui"
)

var Commands map[string]cli.CommandFactory
var Ui ui.Ui

func init() {
	Commands = map[string]cli.CommandFactory{
		"serve": func() (cli.Command, error) {
			return &commands.Serve{Ui}, nil
		},
	}
}
