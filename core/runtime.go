package core

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/Popoola-Opeyemi/meeseeks/util"
	execute "github.com/alexellis/go-execute/pkg/v1"
	"go.uber.org/zap"
)

var settingFN string = "config.json"

// const ...
const (
	Linux   string = "linux"
	Windows        = "windows"
)

// Instance ...
type Instance struct {
	Logger          *zap.SugaredLogger
	Config          *Commands
	OperatingSystem string
}

// OperationStatus ...
type OperationStatus struct {
	Status   int
	Result   execute.ExecResult
	Duration time.Duration
}

// InitApplication starts the application by initializing the config and logger ...
func InitApplication() (i Instance, e error) {

	// starting logger
	logger := util.InitLogger().Sugar()

	// ensures logger closes when done
	defer logger.Sync()

	i.Logger = logger

	i.OperatingSystem = runtime.GOOS

	logger.Info(splashScreen())

	logger.Info("Starting Application \n")

	// checking to ensure file exists
	fExist := FileExist(settingFN)

	// if file does not exist attempt to create it
	if !fExist {
		fmt.Print("Cannot Find %s , creating %s \n", settingFN, settingFN)

		err := CreateConfig(settingFN)

		if err != nil {
			logger.Info("Cannot create file exiting ... \n")
			os.Exit(1)
		}
	}

	logger.Info("Reading Config File ... \n")

	config, err := ReadConfig(settingFN)

	if err != nil {
		logger.Info("[Exiting !] cannot read Config file \n")
		os.Exit(1)
	}

	i.Config = config

	return

}

// RunCommand runs command using the passed function parameters
func RunCommand(command string, directory string, logger *zap.SugaredLogger, result chan OperationStatus) {
	start := time.Now()

	logger.Infof("+c [running cmd] %s \n", command)

	cmd := execute.ExecTask{
		Command: command,
		Cwd:     directory,
	}

	res, err := cmd.Execute()

	end := time.Since(start)

	if err != nil {
		res.Stderr = err.Error()
		// initializing and populating OperationStatus struct
		s := OperationStatus{Status: Error, Result: res, Duration: end}
		result <- s
	}

	if res.ExitCode != 0 {
		s := OperationStatus{Status: Failed, Result: res, Duration: end}
		// sending s to result channel
		result <- s
	}

	if res.ExitCode == 0 {
		s := OperationStatus{Status: Success, Result: res, Duration: end}
		result <- s
	}

}

// RunCommandSync runs command synchronously ...
func RunCommandSync(command string, directory string, logger *zap.SugaredLogger) (result OperationStatus) {

	// get current time to track execution of task
	start := time.Now()

	logger.Infof("+s [running cmd]  %s \n", command)

	cmd := execute.ExecTask{
		Command: command,
		Cwd:     directory,
	}

	res, err := cmd.Execute()

	// get time elapsed time between start and end
	end := time.Since(start)

	if err != nil {
		result.Status = Error
		result.Result = res

		result.Duration = end
	}
	if res.ExitCode != 0 {
		result.Status = Failed
		result.Result = res
		result.Duration = end
	}

	if res.ExitCode == 0 {
		result.Status = Success
		result.Result = res
		result.Duration = end
	}

	start = time.Now()
	return result

}

func splashScreen() string {
	splash := `
___________________________________________
___________________________________________
|  \/  | ___  ___  ___  ___  ___| | _____ 
| |\/| |/ _ \/ _ \/ __|/ _ \/ _ \ |/ / __|
| |  | |  __/  __/\__ \  __/  __/   <\__ \
|_|  |_|\___|\___||___/\___|\___|_|\_\___/
___________________________________________
___________________________________________                                   
`
	return splash
}
