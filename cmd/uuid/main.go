package main

import "github.com/exograd/go-program"

func main() {
	p := program.NewProgram("uuid", "utilities for the go-uuid library")

	p.SetMain(mainCmd)

	p.ParseCommandLine()
	p.Run()
}

func mainCmd(p *program.Program) {
	// TODO
}
