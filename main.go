package main

import (
	"os"

	"github.com/mateeullahmalik/goa-demo/cmd"
	"github.com/mateeullahmalik/goa-demo/internal/errors"
	"github.com/mateeullahmalik/goa-demo/internal/log"
	"github.com/mateeullahmalik/goa-demo/internal/sys"
)

const (
	debugModeEnvName = "SOLO_SERVICE_DEBUG"
)

var (
	debugMode = sys.GetBoolEnv(debugModeEnvName, false)
)

func main() {
	defer errors.Recover(log.FatalAndExit)

	app := cmd.NewApp()
	err := app.Run(os.Args)

	log.FatalAndExit(err)
}

func init() {
	log.SetDebugMode(debugMode)
}
