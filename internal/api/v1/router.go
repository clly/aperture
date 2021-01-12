package v1

import (
	"clly/apterture/internal/api/v1/helloworld"
	"clly/apterture/internal/api/v1/stateapi"
	"clly/apterture/pkg/storage"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(db storage.StateStorage) http.Handler {
	mixer := mux.NewRouter()
	helo := helloworld.HelloWorld()
	stateAPI := stateapi.NewStateAPI(db)

	mixer.Path("/api/v1/helloworld").Methods(http.MethodGet).Handler(helo)
	mixer.Path("/api/v1/helloworld").Methods(http.MethodPost).Handler(helo)
	mixer.Path("/api/v1/tf").Methods(http.MethodGet).Handler(stateAPI)
	mixer.Path("/api/v1/tf").Methods(http.MethodPost).Handler(stateAPI)
	return mixer
}
