package main

import (
	loggerPkg "github.com/lapeko/orca-test/internal/api/logger"
	"os"
)

func getJsonInputPath() string {
	logger := loggerPkg.GetInstance()

	if len(os.Args) < 2 {
		logger.Error("no provided json input path from arguments")
		os.Exit(1)
	}

	return os.Args[1]
}
