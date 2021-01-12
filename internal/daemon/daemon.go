package daemon

import (
	v1 "clly/apterture/internal/api/v1"
	"net/http"

	"github.com/gorilla/mux"
)

type Aperture struct {
	router *mux.Router
}

type initers func(monitor *Aperture) error

func (d *Aperture) init() error {
	for _, initer := range []initers{
		initMux,
	} {
		if err := initer(d); err != nil {
			return err
		}
	}
	return nil
}

func initMux(d *Aperture) error {
	d.router = mux.NewRouter()

	d.router.PathPrefix("/api/v1").Handler(v1.Handler())
	return nil
}

func (d *Aperture) Run() error {
	if err := d.init(); err != nil {
		return err
	}
	addr := ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: d.router,
	}
	return server.ListenAndServe()
}
