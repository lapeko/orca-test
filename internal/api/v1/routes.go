package v1

import (
	"fmt"
	"github.com/lapeko/orca-test/internal/api/statistics"
	"github.com/lapeko/orca-test/internal/vulnerability"
	"net/http"
)

const pathPrefix = "/api/v1"

var services = struct {
	Vulnerability vulnerability.Service
	Stats         statistics.Statistics
}{}

func RegisterRoutes(
	vs vulnerability.Service,
	mux *http.ServeMux,
	stats statistics.Statistics,
) {
	services.Stats = stats
	services.Vulnerability = vs

	mux.HandleFunc(fmt.Sprintf("%s/attack", pathPrefix), attackHandler)
	mux.HandleFunc(fmt.Sprintf("%s/stats", pathPrefix), statsHandler)
}
