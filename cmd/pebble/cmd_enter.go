package main

import "github.com/jessevdk/go-flags"

var shortEnterHelp = "Run the pebble environment and optionally execute command"
var longEnterHelp = `
The enter command starts pebble and runs the configured environment.
Optionally, it also executes a command.
`

type cmdEnter struct {
	cmdRun
}

func init() {
	addCommand("enter", shortEnterHelp, longEnterHelp, func() flags.Commander { return &cmdEnter{} },
		runOptionsHelp, nil)
}

func (rcmd *cmdEnter) Execute(args []string) error {
	if len(args) == 0 {
		args = nil
	}
	rcmd.run(args)

	return nil
}
