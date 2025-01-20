package main

import (
	"github.com/lapeko/orca-test/internal/api"
	loggerPkg "github.com/lapeko/orca-test/internal/api/logger"
	"os"
)

const port = ":80"

func main() {
	logger := loggerPkg.GetInstance()

	a := api.NewApi()
	a.Init(getJsonInputPath())
	a.SetupRoutes()

	if err := a.Listen(port); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
