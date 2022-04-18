package handler

import (
	"net/http"

	"github.com/DionisiyGri/sample-go/internal/db"
	log "github.com/sirupsen/logrus"
)

// Health is typical handler to check health status of app. It checks db connection.
// --also you can add connection to other services like queues, other dbs etc.
func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	if err := db.Ping(&h.Config); err != nil {
		log.Errorf("database connection failure: %s", err.Error())
		writeErrResponse(w, http.StatusInternalServerError, "")
		return
	}

	writeOkResponse(w, "The app is healthy")
}
