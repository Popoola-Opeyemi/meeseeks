package core

import (
	"sync"

	execute "github.com/alexellis/go-execute/pkg/v1"
	"go.uber.org/zap"
)

func (i *Instance) handleConcurrent(status chan int) {
	var wg sync.WaitGroup

	// getting logger the pointer
	logger := i.Logger

	cmdList := i.Config.Commands.List

	// looping through each of the list and for selected commands
	for idx, cmd := range cmdList {

		wg.Add(idx)

		// concurrent handler
		if cmd.Concurrent {

			for i, list := range cmd.List {
				_ = i
				go runCommand(list.CMD, cmd.Directory, logger)
			}

		}

		if !cmd.Concurrent {
			for i, list := range cmd.List {
				_ = i
				runCommand(list.CMD, cmd.Directory, logger)
			}
		}

	}
}

func runCommand(command string, directory string, logger *zap.SugaredLogger) (int, execute.ExecResult, error) {

	logger.Debug("running command %s", command)

	cmd := execute.ExecTask{
		Command: command,
		Cwd:     directory,
	}

	res, err := cmd.Execute()

	if err != nil {
		return Failed, res, err

	}

	if res.ExitCode != 0 {
		return Failed, res, nil
	}

	return Success, res, nil
}

// Handler ...
func (i *Instance) StartHandler() {

	buffChan := make(chan int, 0)
	// assigning the config to c to use for comparision
	c := i.Config

	// calls the function to run task concurrently
	// starting up a go routine
	if c.Commands.Concurrent {
		go i.handleConcurrent(buffChan)
	}

	// calls handler function to run task synchronously
	if c.Commands.Concurrent == false {
		i.handleConcurrent(buffChan)
	}

}
