package api

import (
	"fmt"
	"github.com/lapeko/orca-test/internal/api/logger"
	"github.com/lapeko/orca-test/internal/api/middlewares"
	"github.com/lapeko/orca-test/internal/api/statistics"
	"github.com/lapeko/orca-test/internal/api/v1"
	"github.com/lapeko/orca-test/internal/cloudparser"
	"github.com/lapeko/orca-test/internal/vulnerability"
	"net/http"
)

type api struct {
	vulnerabilityService vulnerability.Service
	mux                  *http.ServeMux
	stats                statistics.Statistics
	logger               logger.Logger
}

type Api interface {
	Init(JSONPath string)
	SetupRoutes()
	Listen(port string) error
}

func (a *api) Init(JSONPath string) {
	a.logger.Info("API initialization...")
	env, err := cloudparser.LoadCloudEnvironment(JSONPath)
	if err != nil {
		panic(err)
	}
	a.vulnerabilityService = vulnerability.NewCalculator(env)
	a.mux = http.NewServeMux()
	a.stats = statistics.NewApiStatistics(len(env.VMs))
	a.logger.Info("API initialized")
}

func (a *api) SetupRoutes() {
	a.logger.Info("API routes initialization...")
	if a.mux == nil {
		panic("API not initialized")
	}
	v1.RegisterRoutes(a.vulnerabilityService, a.mux, a.stats)
	a.logger.Info("API routes initialized")
}

func (a *api) Listen(port string) error {
	handlerWithMiddleware := middlewares.RequestLogger(a.mux)
	a.logger.Info(fmt.Sprintf("Server is working on port %s", port))
	return http.ListenAndServe(port, handlerWithMiddleware)
}

func NewApi() Api {
	log := logger.GetInstance()
	return &api{logger: log}
}
