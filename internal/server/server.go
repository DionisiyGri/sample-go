package server

import (
	"net/http"

	"github.com/DionisiyGri/sample-go/internal/config"
	"github.com/DionisiyGri/sample-go/internal/router"
	"github.com/DionisiyGri/sample-go/internal/sample"
	log "github.com/sirupsen/logrus"
)

// Init starts new server at provided port
func Init(port string, app *sample.App, c config.Config) error {
	r := router.Init(app, c)
	log.WithField("port", port).Debugf("App starting at port=%s", port)
	http.Handle("/", r)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Errorf("error  server Init: %s", err)
		return err
	}

	return nil
}
