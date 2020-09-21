package core

import (
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

type OperationStatus struct {
	Status   int
	Result   execute.ExecResult
	Duration time.Duration
}

// InitInstance ...
func InitApplication() (i Instance, e error) {

	logger := util.InitLogger().Sugar()

	i.Logger = logger

	i.OperatingSystem = runtime.GOOS

	config, err := ReadSettings(settingFN)

	if err != nil {
		logger.Debug("an error ocurred", err)
		return i, e
	}

	i.Config = config

	return i, e

}

// RunCommand runs command using the passed function parameters
func RunCommand(command string, directory string, logger *zap.SugaredLogger, result chan OperationStatus) {
	start := time.Now()

	logger.Debugf("running command %s", command)

	cmd := execute.ExecTask{
		Command: command,
		Cwd:     directory,
	}

	res, err := cmd.Execute()

	end := time.Since(start)

	if err != nil {
		s := OperationStatus{Status: Error, Result: res, Duration: end}
		result <- s
		return
	}

	if res.ExitCode != 0 {
		s := OperationStatus{Status: Failed, Result: res, Duration: end}
		result <- s
		return
	}

	if res.ExitCode == 0 {
		s := OperationStatus{Status: Success, Result: res, Duration: end}
		result <- s

	}

}

// RunCommand runs command using the passed function parameters
func RunCommandSync(command string, directory string, logger *zap.SugaredLogger) (result OperationStatus) {
	start := time.Now()

	logger.Debugf("[running cmd]  %s \n", command)

	cmd := execute.ExecTask{
		Command: command,
		Cwd:     directory,
	}

	res, err := cmd.Execute()
	end := time.Since(start)

	if err != nil {
		result.Status = Error
		result.Result = res
		result.Duration = end
		return
	}

	if res.ExitCode != 0 {
		result.Status = Failed
		result.Result = res
		result.Duration = end
		return
	}

	if res.ExitCode == 0 {
		result.Status = Success
		result.Result = res
		result.Duration = end
		return
	}

	return result

}
