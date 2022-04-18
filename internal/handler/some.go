package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DionisiyGri/sample-go/internal/model"
)

// GetAllSome handler to return some array of objects.
func (h *Handlers) GetAllSome(w http.ResponseWriter, r *http.Request) {
	items, err := h.app.GetAll()
	if err != nil {
		errors := map[string]string{"db": err.Error()}
		writeErrResponse(w, http.StatusInternalServerError, errors)
		return
	}

	writeListResponse(w, items)
}

// PostSome handler to add new entity.
func (h *Handlers) PostSome(w http.ResponseWriter, r *http.Request) {
	someEntity := &model.Some{}
	err := json.NewDecoder(r.Body).Decode(someEntity)
	if err != nil {
		errors := map[string]string{"parsing": err.Error()}
		writeErrResponse(w, http.StatusBadRequest, errors)
		return
	}

	if !h.validate(someEntity, w) {
		return
	}

	item, err := h.app.PostSome(someEntity)
	if err != nil {
		errors := map[string]string{"db": err.Error()}
		writeErrResponse(w, http.StatusInternalServerError, errors)
		return
	}

	writeOkResponse(w, item)
}
