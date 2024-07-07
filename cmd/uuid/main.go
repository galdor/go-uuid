package main

import (
	"fmt"
	"math"
	"strconv"

	"go.n16f.net/program"
	"go.n16f.net/uuid"
)

func main() {
	p := program.NewProgram("uuid", "utilities for the go-uuid library")

	p.AddOption("v", "version", "version", "4",
		"the version of the uuid to generate")

	p.SetMain(mainCmd)

	p.ParseCommandLine()
	p.Run()
}

func mainCmd(p *program.Program) {
	vs := p.OptionValue("version")
	i64, err := strconv.ParseInt(vs, 10, 64)
	if err != nil || i64 <= 0 || i64 > math.MaxInt32 {
		p.Fatal("invalid version %q", vs)
	}
	v := uuid.Version(i64)

	var id uuid.UUID
	if err := id.Generate(v); err != nil {
		p.Fatal("cannot generate uuid: %v", err)
	}

	fmt.Println(id.String())
}
