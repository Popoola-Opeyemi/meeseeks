package main

import (
	"github.com/Popoola-Opeyemi/meeseeks/util"
)

func main() {
	logger := util.InitLogger().Sugar()
	defer logger.Sync()
	logger.Debugf("%s", util.IsLinux())
}
