package v1

import (
	"fmt"
	loggerPkg "github.com/lapeko/orca-test/internal/api/logger"
	"net/http"
	"strings"
)

func GetApiRequestInfo(r *http.Request) string {
	queryParams := r.URL.Query().Encode()
	if queryParams != "" {
		queryParams = fmt.Sprintf("?%s", queryParams)
	}

	return fmt.Sprintf("%s %s%s", strings.ToUpper(r.Method), r.URL.Path, queryParams)
}

func sendError(w http.ResponseWriter, r *http.Request, error string, code int) {
	logger := loggerPkg.GetInstance()
	logger.Error(fmt.Sprintf("%s %d\n\t%s", GetApiRequestInfo(r), code, error))
	http.Error(w, error, code)
}

func sendSuccess(w http.ResponseWriter, r *http.Request, data []byte) {
	logger := loggerPkg.GetInstance()
	code := http.StatusOK
	logger.Info(fmt.Sprintf("%s %d", GetApiRequestInfo(r), code))
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
