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

	// if operation is to be ran concurrently
	if c.Concurrent {
		for _, list := range c.List {
			if list.CMD != "" {
				wg.Add(1)
				go startConcurrent(list.CMD, c.Directory, l)
			}
		}
	}

	// if operation is to be ran synchronously
	if !c.Concurrent {
		for _, list := range c.List {
			if list.CMD != "" {
				startSynchronous(list.CMD, c.Directory, l)
			}
		}
	}
}

// Sync ...
func (h HandlerObjects) Sync() {
	c := h.Config
	l := h.Logger

	// if operation is to be ran concurrently
	if c.Concurrent {

		wg.Add(len(c.List))

		for _, list := range c.List {

			// ensuring cmd is not empty
			if list.CMD != "" {

				// adds to waitgroup counter

				// startup the run command
				go startConcurrent(list.CMD, c.Directory, l)

			}
		}
		// ensuring the main thread waits until the goroutine has finished
		wg.Wait()

	}

	if !c.Concurrent {
		for _, list := range c.List {

			// ensuring cmd is not empty
			if list.CMD != "" {

				// calling startSynchronous handler
				startSynchronous(list.CMD, c.Directory, l)
			}
		}
	}
}

func startConcurrent(command string, directory string, logger *zap.SugaredLogger) {

	// ensuring the wait group sends done once function finished
	defer wg.Done()

	// creating a channel
	resChan := make(chan OperationStatus)

	defer close(resChan)

	// starting go routine to handle running command
	go RunCommand(command, directory, logger, resChan)

	// sending result to channel
	result := <-resChan

	// checking the result returned on channel
	switch result.Status {

	case Failed:
		logger.Errorf("[Failed] command %s \n", command)

	case Success:
		logger.Infof("[success] command %s finished in [%s] \n", command, result.Duration)

	case Error:
		logger.Errorf("[Error] command %s encountered errors \n", command)

	}

}

// starts an asynchronous handler to run the command
func startSynchronous(command string, directory string, logger *zap.SugaredLogger) {

	result := RunCommandSync(command, directory, logger)

	switch result.Status {

	case Failed:
		logger.Errorf("[Failed] command %s \n", command)

	case Success:
		logger.Infof("[success] command %s finished in [%s] \n", command, result.Duration)

	case Error:
		logger.Errorf("[Error] command %s encountered errors \n", command)

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

		wg.Add(len(cmdList))
		for _, cmd := range cmdList {

			// increment wait group by 1

			h = HandlerObjects{Logger: logger, Config: cmd}

			// calls the concurrent method
			go h.Concurrent()
		}

		// wait for the execution of all goroutines
		wg.Wait()

	}

	// synchronous operations
	if !concurrent {

		for _, cmd := range cmdList {

			h = HandlerObjects{Logger: logger, Config: cmd}
			// calls sync method
			h.Sync()
		}
	}

}
