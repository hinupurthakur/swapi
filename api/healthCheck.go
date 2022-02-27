package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Infoln("GET: healthcheck API")
	result := `{"status": "Connection successful"}`
	var output map[string]interface{}
	err := json.Unmarshal([]byte(result), &output)
	if err != nil {
		log.Errorln("healthCheck: unable to encode string to JSON", err)
		http.Error(w, "healthCheck: unable to encode string to JSON", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		log.Errorln("healthCheck: unable to encode output", err)
		http.Error(w, "healthCheck: unable to encode output", http.StatusBadRequest)
		return
	}
}
