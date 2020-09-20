package core

import (
	"runtime"

	"github.com/Popoola-Opeyemi/meeseeks/util"
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
