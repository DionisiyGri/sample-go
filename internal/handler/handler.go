package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DionisiyGri/sample-go/internal/config"
	"github.com/DionisiyGri/sample-go/internal/model"
	"github.com/DionisiyGri/sample-go/internal/sample"
	"github.com/DionisiyGri/sample-go/internal/validation"
	log "github.com/sirupsen/logrus"
)

// Handlers provides https handlers
type Handlers struct {
	app    *sample.App
	Config config.Config
}

// NewHandlers initialize handler
func NewHandlers(app *sample.App, c config.Config) *Handlers {
	return &Handlers{
		app:    app,
		Config: c,
	}
}

// validate validates struct and writes errorn if it's not valid
func (h *Handlers) validate(obj validation.Validatable, w http.ResponseWriter) bool {
	errors := validation.CustomValidateStruct(obj)

	if len(errors) != 0 {
		writeErrResponse(w, http.StatusBadRequest, errors)
		return false
	}

	return true
}

// writeOkResponse generic func to return single object
func writeOkResponse(w http.ResponseWriter, data interface{}) {
	response := model.APIResponse{}
	response.Type = "item"
	response.Data = data

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.WithField("method", "writeOkResponse").Errorf("Encode error: %s", err.Error())
	}
}

// writeErrResponse generic func to return error with passed error code
func writeErrResponse(w http.ResponseWriter, code int, data interface{}) {
	response := model.APIResponse{}
	response.Type = "request_error"

	response.Data = model.APIErrorResponse{Errors: data}

	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.WithField("method", "writeErrResponse").Errorf("Encode error: %s", err.Error())
	}
}

// writeListResponse generic func to return array of data
func writeListResponse(w http.ResponseWriter, data interface{}) {
	response := model.APIResponse{}
	response.Type = "list"
	response.Data = data

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.WithField("method", "writeListResponse").Errorf("Encode error: %s", err.Error())
	}
}
