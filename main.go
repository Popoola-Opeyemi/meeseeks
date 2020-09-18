package main

import (
	"automator/util"
)

func main() {
	logger := util.InitLogger().Sugar()
	defer logger.Sync()
	logger.Debugf("%s", util.IsLinux())
}
