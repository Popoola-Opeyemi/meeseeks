package main

import (
	"os"

	"github.com/Popoola-Opeyemi/meeseeks/core"
)

func main() {

	instance, err := core.InitApplication()

	if err != nil {
		os.Exit(1)
		return
	}

	instance.StartHandler()

	instance.Logger.Info("All Operation Finished")

}
