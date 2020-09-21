package main

import (
	"os"
	"time"

	"github.com/Popoola-Opeyemi/meeseeks/core"
)

func main() {
	start := time.Now()

	instance, err := core.InitApplication()

	if err != nil {
		os.Exit(1)
		return
	}

	instance.StartHandler()

	end := time.Since(start)

	instance.Logger.Infof("Operation Completed in [%s]", end)

}
