package main

import "io"

type handler func(in []string, out io.Writer) (err error)

var (
	registers = map[string]handler{

		"interface": InterfaceHandler,
	}
)
