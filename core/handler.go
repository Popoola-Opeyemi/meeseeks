package core

import (
	"os/exec"
	"sync"
)

func (i *Instance) handleConcurrent() {
	var wg sync.WaitGroup

	// getting logger the pointer
	logger := i.Logger

	cmdList := i.Config.Commands.List

	// looping through each of the list and for selected commands
	for idx, cmd := range cmdList {
		wg.Add(idx)
		if cmd.Concurrent {
			go runConcurrentCommand(cmd.List)
		}

		runCommand(cmd.List)
	}

	_ = logger

}

func runCommand([]ListItems) {

}

func runConcurrentCommand([]ListItems) {

}

func runCmd(cmd string) (CommandOut, error) {

	output := CommandOut{}

	cmdRunner := exec.Command(cmd)

	stdout, err := cmdRunner.Output()

	if err != nil {
		return output, err
	}

	output.StdOutput = stdout

	stderr, err := cmdRunner.CombinedOutput()

	if err != nil {
		return output, err
	}

	output.StdError = stderr

	return output, nil

}

// Handler ...
func (i *Instance) StartHandler() {

	var wg sync.WaitGroup

	// assigning the config to c to use for comparision
	c := i.Config

	// calls the function to run task concurrently
	// starting up a go routine
	if c.Commands.Concurrent {
		go i.handleConcurrent()
	}

	// calls handler function to run task synchronously
	if c.Commands.Concurrent == false {
		i.handleConcurrent()
	}

	wg.Wait()
}
