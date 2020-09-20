package main

import (
	"os"

	"github.com/Popoola-Opeyemi/meeseeks/core"
)

func main() {

	instance, err := core.InitApplication()

	if err != nil {
		os.Exit(1)
	}

	instance.StartHandler()

}
