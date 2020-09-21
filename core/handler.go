package core

import (
	"sync"

	execute "github.com/alexellis/go-execute/pkg/v1"
	"go.uber.org/zap"
)

// handleConcurrent function loops loops through each command to determine
// which inner go routine to be spawned
func handleConcurrent(items JsonInner, wg *sync.WaitGroup, logger *zap.SugaredLogger) {

	if items.Concurrent {
		for i, list := range items.List {
			_ = i
			wg.Add(1)
			go runCommand(list.CMD, items.Directory, wg, logger)
		}
	}

	if !items.Concurrent {
		for i, list := range items.List {
			_ = i
			runCommand(list.CMD, items.Directory, wg, logger)
		}
	}

	defer wg.Done()

}

// runCommand runs command using the passed function parameters
func runCommand(command string, directory string, wg *sync.WaitGroup, logger *zap.SugaredLogger) {

	logger.Debugf("running command %s", command)

	cmd := execute.ExecTask{
		Command: command,
		Cwd:     directory,
	}

	res, err := cmd.Execute()

	if err != nil {
		logger.Debug("an error occured", err)
	}

	if res.ExitCode == 0 {
		logger.Debugf("command %s is finished", command)
	}

	defer wg.Done()

	// if err != nil {
	// 	return Failed, res, err

	// }

	// if res.ExitCode != 0 {
	// 	return Failed, res, nil
	// }

	// return Success, res, nil
}

// StartHandler function responsible for starting appropriate ...
func (i *Instance) StartHandler() {

	wg := new(sync.WaitGroup)

	// assigning the commands list
	cmdList := i.Config.Commands.List

	// assigning concurrent status
	concurrent := i.Config.Commands.Concurrent

	// assigning the logger
	logger := i.Logger

	// concurrent operations
	if concurrent {

		for idx, cmd := range cmdList {
			_ = idx
			wg.Add(1)
			// calls the function to run task concurrently
			go handleConcurrent(cmd, wg, logger)
		}

		wg.Wait()

	}

	// // synchronous operations
	// if !concurrent {

	// 	for idx, cmd := range cmdList {
	// 		_ = idx

	// 		// calls the function to run task synchronously
	// 		handleConcurrent(cmd, logger)
	// 	}
	// }

}
