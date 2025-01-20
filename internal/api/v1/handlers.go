package v1

import (
	"encoding/json"
	"net/http"
	"time"
)

func attackHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vmId := r.URL.Query().Get("vm_id")

	if vmId == "" {
		sendError(w, r, "Missing \"vm_id\" query parameter", http.StatusBadRequest)
		return
	}

	threats, err := services.Vulnerability.ScanVMThreads(vmId)

	if err != nil {
		sendError(w, r, err.Error(), http.StatusNotFound)
		return
	}

	data, err := json.Marshal(threats)
	if err != nil {
		sendError(w, r, "failed to encode response", http.StatusInternalServerError)
		return
	}

	services.Stats.AddStats(time.Since(start).Nanoseconds())
	sendSuccess(w, r, data)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats := services.Stats.GetStats()

	data, err := json.Marshal(stats)
	if err != nil {
		sendError(w, r, "failed to encode response", http.StatusInternalServerError)
		return
	}

	sendSuccess(w, r, data)
}
