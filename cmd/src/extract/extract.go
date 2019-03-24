package main

import (
	"fmt"
	"os"
)

func main() {
	taskName := os.Args[1]

	task := registers[taskName]

	err := task(os.Args[1:], os.Stdout)

	fmt.Println(err)
}
