package core

import (
	"sync"

	"go.uber.org/zap"
)

var wg = new(sync.WaitGroup)

// Concurrent ...
func (h HandlerObjects) Concurrent() {
	defer wg.Done()

	c := h.Config
	l := h.Logger

	if c.Concurrent {
		for _, list := range c.List {
			wg.Add(1)
			go startConcurrent(list.CMD, c.Directory, l)
		}
	}

	if !c.Concurrent {
		for _, list := range c.List {
			startSynchronous(list.CMD, c.Directory, l)
		}
	}
}

// Sync ...
func (h HandlerObjects) Sync() {
	c := h.Config
	l := h.Logger

	if c.Concurrent {
		for _, list := range c.List {

			// adds to concurrent list
			wg.Add(1)

			// startup the run command
			go startConcurrent(list.CMD, c.Directory, l)
		}

	}

	if !c.Concurrent {
		for _, list := range c.List {
			startSynchronous(list.CMD, c.Directory, l)
		}
	}
}

func startConcurrent(command string, directory string, logger *zap.SugaredLogger) {
	defer wg.Done()

	resChan := make(chan OperationStatus)

	go RunCommand(command, directory, logger, resChan)

	result := <-resChan

	switch result.Status {

	case Failed:
		logger.Errorf("[Failed] command %s \n", command)

	case Success:
		logger.Debugf("[success] command %s \n", command)

	case Error:
		logger.Errorf("[Error] command %s encountered errors \n", command)

	}

}

func startSynchronous(command string, directory string, logger *zap.SugaredLogger) {

	result := RunCommandSync(command, directory, logger)

	switch result.Status {

	case Failed:
		logger.Debug("Operation Failed")

	case Success:
		logger.Debug("Operaton finished successfully")

	case Error:
		logger.Debug("Operaton encountered Errors")

	}

}

// StartHandler function responsible for starting appropriate ...
func (i *Instance) StartHandler() {

	// assigning the commands list
	cmdList := i.Config.Commands.List

	// assigning concurrent status
	concurrent := i.Config.Commands.Concurrent

	// assigning the logger
	logger := i.Logger

	var h HandlerObjects

	// concurrent operations
	if concurrent {

		for _, cmd := range cmdList {

			// increment wait group by 1
			wg.Add(1)

			h = HandlerObjects{Logger: logger, Config: cmd}
			// calls the function to run task concurrently
			go h.Concurrent()
		}

		// wait for the execution of all goroutines
		wg.Wait()

	}

	// synchronous operations
	if !concurrent {

		for _, cmd := range cmdList {

			h = HandlerObjects{Logger: logger, Config: cmd}
			h.Sync()
		}
	}

}
