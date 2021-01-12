package v1

import (
	"clly/apterture/internal/api/v1/helloworld"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	mixer := mux.NewRouter()
	helo := helloworld.HelloWorld()

	mixer.Path("/api/v1/helloworld").Methods(http.MethodGet).Handler(helo)
	mixer.Path("/api/v1/helloworld").Methods(http.MethodPost).Handler(helo)
	return mixer
}
