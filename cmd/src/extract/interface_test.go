package main

import (
	"os"
	"testing"
)

type A struct{}

func (*A) M1() {}

func Test_Interface(t *testing.T) {

	err := InterfaceHandler([]string{"interface_test.go"}, os.Stdout)

	println(err)
}
