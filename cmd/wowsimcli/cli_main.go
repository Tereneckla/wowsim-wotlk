package main

import (
	"github.com/Tereneckla/wowsim-wotlk/cmd/wowsimcli/cmd"
	"github.com/Tereneckla/wowsim-wotlk/sim"
)

func init() {
	sim.RegisterAll()
}

// Version information.
// This variable is set by the makefile in the release process.
var Version string

func main() {
	cmd.Execute(Version)
}
