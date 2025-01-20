package middlewares

import (
	loggerPkg "github.com/lapeko/orca-test/internal/api/logger"
	"github.com/lapeko/orca-test/internal/api/v1"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	logger := loggerPkg.GetInstance()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(v1.GetApiRequestInfo(r))

		next.ServeHTTP(w, r)
	})
}
