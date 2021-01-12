package daemon

import (
	v1 "clly/apterture/internal/api/v1"
	"clly/apterture/pkg/storage"
	"net/http"

	"github.com/gorilla/mux"
)

type Aperture struct {
	router *mux.Router
	db storage.StateStorage
}

type initers func(aperture *Aperture) error

func (d *Aperture) init() error {
	for _, initer := range []initers{
		initStorage,
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

	d.router.PathPrefix("/api/v1").Handler(v1.Handler(d.db))
	return nil
}

func initStorage(d *Aperture) error {
	d.db = storage.NewMemStorage()
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
