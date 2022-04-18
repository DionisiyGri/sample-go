package router

import (
	"github.com/DionisiyGri/sample-go/internal/config"
	"github.com/DionisiyGri/sample-go/internal/handler"
	"github.com/DionisiyGri/sample-go/internal/sample"
	"github.com/gorilla/mux"
)

// Init return router
func Init(app *sample.App, c config.Config) *mux.Router {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/" + c.ServiceName).Subrouter()

	h := handler.NewHandlers(app, c)

	router.HandleFunc("/all", h.GetAllSome).Methods("GET")
	router.HandleFunc("/all", h.PostSome).Methods("POST")
	router.HandleFunc("/health", h.Health).Methods("GET")

	return router
}
